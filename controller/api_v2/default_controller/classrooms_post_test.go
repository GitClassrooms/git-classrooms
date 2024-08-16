package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.hs-flensburg.de/gitlab-classroom/config"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database/query"
	gitlabRepoMock "gitlab.hs-flensburg.de/gitlab-classroom/repository/gitlab/_mock"
	"gitlab.hs-flensburg.de/gitlab-classroom/repository/gitlab/model"
	mailRepoMock "gitlab.hs-flensburg.de/gitlab-classroom/repository/mail/_mock"
	"gitlab.hs-flensburg.de/gitlab-classroom/utils"
	"gitlab.hs-flensburg.de/gitlab-classroom/utils/factory"
	fiberContext "gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/session"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateClassroom(t *testing.T) {
	// setup database
	restoreDatabase(t)
	db, err := gorm.Open(postgres.Open(integrationTest.dbURL))
	if err != nil {
		t.Fatal(err)
	}

	query.SetDefault(db)

	user := factory.User()

	gitlabRepo := gitlabRepoMock.NewMockRepository(t)
	mailRepo := mailRepoMock.NewMockRepository(t)

	app := fiber.New()
	app.Use("/api", func(c *fiber.Ctx) error {
		ctx := fiberContext.Get(c)
		ctx.SetGitlabRepository(gitlabRepo)
		ctx.SetUserID(user.ID)

		s := session.Get(c)
		s.SetUserState(session.LoggedIn)
		s.SetUserID(user.ID)

		s.Save()
		return c.Next()
	})

	handler := NewApiV2Controller(mailRepo, config.ApplicationConfig{})

	t.Run("CreateClassroom", func(t *testing.T) {
		app.Post("/api/v2/classrooms", handler.CreateClassroom)

		requestBody := createClassroomRequest{
			Name:                    gofakeit.Name(),
			Description:             gofakeit.Dessert(),
			CreateTeams:             utils.NewPtr(false),
			MaxTeams:                utils.NewPtr(0),
			MaxTeamSize:             1,
			StudentsViewAllProjects: utils.NewPtr(false),
		}

		gitlabRepo.
			EXPECT().
			CreateGroup(
				requestBody.Name,
				model.Private,
				requestBody.Description,
			).
			Return(
				&model.Group{ID: 1},
				nil,
			).
			Times(1)

		gitlabRepo.
			EXPECT().
			CreateGroupAccessToken(
				1,
				"Gitlab Classrooms",
				model.OwnerPermissions,
				mock.AnythingOfType("time.Time"),
				"api",
			).
			Return(
				&model.GroupAccessToken{ID: 20, Token: "token"},
				nil,
			).
			Times(1)

		req := newPostJsonRequest("/api/v2/classrooms", requestBody)
		resp, err := app.Test(req)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
		assert.NoError(t, err)

		classRoom, err := query.Classroom.
			WithContext(context.Background()).
			Where(query.Classroom.OwnerID.Eq(user.ID)).
			First()

		assert.NoError(t, err)
		assert.Equal(t, requestBody.Name, classRoom.Name)
		assert.Equal(t, requestBody.Description, classRoom.Description)
		assert.Equal(t, 1, classRoom.GroupID)
		assert.Equal(t, 20, classRoom.GroupAccessTokenID)
		assert.Equal(t, "token", classRoom.GroupAccessToken)
		assert.Equal(t, false, classRoom.StudentsViewAllProjects)

		assert.Equal(t, fmt.Sprintf("/api/v2/classrooms/%s", classRoom.ID.String()), resp.Header.Get("Location"))
	})
}
