package default_controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.hs-flensburg.de/gitlab-classroom/context"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database/query"
)

type getMeClassroomAssignmentsResponse struct {
	database.AssignmentProjects
	ProjectPath string `json:"projectPath"`
}

func (ctrl *DefaultController) GetMeClassroomAssignments(c *fiber.Ctx) error {
	classroom := context.GetClassroom(c)

	queryAssignment := query.Assignment
	queryAssignmentProjects := query.AssignmentProjects
	assignmentProjects, err := queryAssignmentProjects.WithContext(c.Context()).
		Join(queryAssignment, queryAssignment.ID.EqCol(queryAssignmentProjects.AssignmentID)).
		Where(queryAssignment.ClassroomID.Eq(classroom.ClassroomID)).
		Find()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	repo := context.GetGitlabRepository(c)
	responses := make([]*getMeClassroomAssignmentsResponse, len(assignmentProjects))
	for i, project := range assignmentProjects {
		webURL := ""
		if project.AssignmentAccepted {
			projectFromGitLab, err := repo.GetProjectById(project.ProjectID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
			webURL = projectFromGitLab.WebUrl
		}
		responses[i] = &getMeClassroomAssignmentsResponse{
			AssignmentProjects: *project,
			ProjectPath:        webURL,
		}
	}

	return c.JSON(responses)
}