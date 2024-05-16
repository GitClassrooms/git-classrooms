package default_controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database/query"
	gitlabRepoMock "gitlab.hs-flensburg.de/gitlab-classroom/repository/gitlab/_mock"
	"gitlab.hs-flensburg.de/gitlab-classroom/repository/gitlab/model"
	mailRepoMock "gitlab.hs-flensburg.de/gitlab-classroom/repository/mail/_mock"
	db_tests "gitlab.hs-flensburg.de/gitlab-classroom/utils/tests"
	contextWrapper "gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/session"
)

func TestPutOwnedAssignments(t *testing.T) {
	testDb := db_tests.NewTestDB(t)

	user := database.User{ID: 1}
	testDb.InsertUser(&user)

	classroom := database.Classroom{
		ID:      uuid.New(),
		OwnerID: user.ID,
	}
	testDb.InsertClassroom(&classroom)
	oldTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	assignment := database.Assignment{
		ID:          uuid.New(),
		Name:        "Test",
		Description: "test",
		DueDate:     &oldTime,
		ClassroomID: classroom.ID,
	}
	testDb.InsertAssignment(&assignment)

	team := database.Team{
		ID:          uuid.New(),
		ClassroomID: classroom.ID,
	}
	testDb.InsertTeam(&team)

	project := database.AssignmentProjects{
		AssignmentID: assignment.ID,
		TeamID:       team.ID,
		ProjectID:    1,
	}
	testDb.InsertAssignmentProjects(&project)

	gitlabRepo := gitlabRepoMock.NewMockRepository(t)
	mailRepo := mailRepoMock.NewMockRepository(t)

	app := fiber.New()
	app.Use("/api", func(c *fiber.Ctx) error {
		ctx := contextWrapper.Get(c)
		ctx.SetOwnedClassroom(&classroom)
		ctx.SetGitlabRepository(gitlabRepo)

		s := session.Get(c)
		s.SetUserState(session.LoggedIn)
		s.SetUserID(user.ID)
		s.Save()
		return c.Next()
	})

	handler := NewApiController(mailRepo)
	app.Put("/api/classrooms/owned/:classroomId/assignments", handler.PutOwnedAssignments)

	targetRoute := fmt.Sprintf("/api/classrooms/owned/%s/assignments/%s", classroom.ID.String(), assignment.ID.String())

	t.Run("updates assignment", func(t *testing.T) {
		newTime := time.Date(2000, 2, 2, 2, 2, 2, 2, time.UTC)
		requestBody := updateAssignmentRequest{
			Name:        "New",
			Description: "new",
			DueDate:     &newTime,
		}

		gitlabRepo.
			EXPECT().
			ChangeProjectName(
				project.ProjectID,
				requestBody.Name,
			).
			Return(
				&model.Project{
					ID:   project.ProjectID,
					Name: requestBody.Name,
				},
				nil,
			).
			Times(1)

		gitlabRepo.
			EXPECT().
			ChangeProjectDescription(
				project.ProjectID,
				requestBody.Description,
			).
			Return(
				&model.Project{
					ID:          project.ProjectID,
					Description: requestBody.Description,
				},
				nil,
			).
			Times(1)

		req := newPutJsonRequest(targetRoute, requestBody)
		resp, err := app.Test(req)

		assert.Equal(t, fiber.StatusAccepted, resp.StatusCode)
		assert.NoError(t, err)

		updatedAssignment, err := query.Assignment.WithContext(context.Background()).Where(query.Assignment.ID.Eq(assignment.ID)).First()
		assert.NoError(t, err)
		assert.Equal(t, requestBody.Name, updatedAssignment.Name)
		assert.Equal(t, requestBody.Description, updatedAssignment.Description)
		assert.Equal(t, newTime, *updatedAssignment.DueDate)

	})

	t.Run("request body is empty", func(t *testing.T) {
		requestBody := updateAssignmentRequest{}

		req := newPutJsonRequest(targetRoute, requestBody)
		resp, err := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
		assert.NoError(t, err)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		bodyString := string(bodyBytes)
		assert.Equal(t, "request body is empty", bodyString)
	})

	t.Run("gitlab error on change project name", func(t *testing.T) {
		requestBody := updateAssignmentRequest{
			Name: "New",
		}

		errorMsg := "gitlab error"
		gitlabRepo.
			EXPECT().
			ChangeProjectName(
				project.ProjectID,
				requestBody.Name,
			).
			Return(
				nil,
				fmt.Errorf(errorMsg),
			).
			Times(1)

		req := newPutJsonRequest(targetRoute, requestBody)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("gitlab error on change project description", func(t *testing.T) {
		requestBody := updateAssignmentRequest{
			Description: "new",
		}

		errorMsg := "gitlab error"
		gitlabRepo.
			EXPECT().
			ChangeProjectDescription(
				project.ProjectID,
				requestBody.Description,
			).
			Return(
				nil,
				fmt.Errorf(errorMsg),
			).
			Times(1)

		req := newPutJsonRequest(targetRoute, requestBody)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

}
