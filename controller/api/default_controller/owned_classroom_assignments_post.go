package default_controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database/query"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
)

type createAssignmentRequest struct {
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	TemplateProjectId int        `json:"templateProjectId"`
	DueDate           *time.Time `json:"dueDate"`
}

func (r createAssignmentRequest) isValid() bool {
	return r.Name != "" && r.TemplateProjectId != 0
}

// @Summary		CreateAssignment
// @Description	CreateAssignment
// @Id				CreateAssignment
// @Tags			assignment
// @Accept			json
// @Param			classroomId		path	string										true	"Classroom ID"	Format(uuid)
// @Param			assignmentInfo	body	default_controller.createAssignmentRequest	true	"Assignment Info"
// @Param			X-Csrf-Token	header	string										true	"Csrf-Token"
// @Success		201
// @Failure		400	{object}	httputil.HTTPError
// @Failure		401	{object}	httputil.HTTPError
// @Failure		403	{object}	httputil.HTTPError
// @Failure		404	{object}	httputil.HTTPError
// @Failure		500	{object}	httputil.HTTPError
// @Router			/classrooms/owned/{classroomId}/assignments [post]
func (ctrl *DefaultController) CreateAssignment(c *fiber.Ctx) error {
	ctx := context.Get(c)
	repo := ctx.GetGitlabRepository()
	classroom := ctx.GetOwnedClassroom()

	requestBody := &createAssignmentRequest{}
	err := c.BodyParser(requestBody)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if !requestBody.isValid() {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// Check if template repository exists
	_, err = repo.GetProjectById(requestBody.TemplateProjectId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Create assigment
	assignmentQuery := query.Assignment
	assignment := &database.Assignment{
		ClassroomID:       classroom.ID,
		TemplateProjectID: requestBody.TemplateProjectId,
		Name:              requestBody.Name,
		Description:       requestBody.Description,
		DueDate:           requestBody.DueDate,
	}

	// Persist assigment
	err = assignmentQuery.WithContext(c.Context()).Create(assignment)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Set("Location", fmt.Sprintf("/api/v1/classrooms/%s/assignments/%s", classroom.ID.String(), assignment.ID.String()))
	return c.SendStatus(fiber.StatusCreated)
}
