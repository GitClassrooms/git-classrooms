package default_controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
)

// @Summary		GetJoinedClassroomAssignments
// @Description	GetJoinedClassroomAssignments
// @Id				GetJoinedClassroomAssignments
// @Tags			assignment
// @Produce		json
// @Param			classroomId	path		string	true	"Classroom ID"	Format(uuid)
// @Success		200			{array}		default_controller.getJoinedClassroomAssignmentResponse
// @Failure		400			{object}	HTTPError
// @Failure		401			{object}	HTTPError
// @Failure		404			{object}	HTTPError
// @Failure		500			{object}	HTTPError
// @Router			/classrooms/joined/{classroomId}/assignment [get]
func (ctrl *DefaultController) GetJoinedClassroomAssignments(c *fiber.Ctx) error {
	ctx := context.Get(c)
	classroom := ctx.GetJoinedClassroom()
	assignments, err := joinedClassroomAssignmentQuery(classroom.ClassroomID, *classroom.TeamID, c).Find()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	repo := ctx.GetGitlabRepository()
	responses := make([]*getJoinedClassroomAssignmentResponse, len(assignments))
	for i, project := range assignments {
		webURL := ""
		if project.AssignmentAccepted {
			projectFromGitLab, err := repo.GetProjectById(project.ProjectID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
			webURL = projectFromGitLab.WebUrl
		}
		responses[i] = &getJoinedClassroomAssignmentResponse{
			AssignmentProjects: project,
			ProjectPath:        webURL,
		}
	}

	return c.JSON(responses)
}
