package default_controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
)

//	@Summary		Get current Team
//	@Description	Get current Team
//	@Tags			team
//	@Accept			json
//	@Produces		json
//	@Param			classroomId	path		string	true	"Classroom ID"	Format(uuid)
//	@Param			teamId		path		string	true	"Team ID"		Format(uuid)
//	@Success		200			{object}	default_controller.getJoinedClassroomTeamResponse
//	@Failure		400			{object}	httputil.HTTPError
//	@Failure		401			{object}	httputil.HTTPError
//	@Failure		404			{object}	httputil.HTTPError
//	@Failure		500			{object}	httputil.HTTPError
//	@Router			/classrooms/joined/{classroomId}/teams/{teamId} [get]
func (ctrl *DefaultController) GetJoinedClassroomTeam(c *fiber.Ctx) error {
	ctx := context.Get(c)

	team := ctx.GetJoinedTeam()

	response := &getJoinedClassroomTeamResponse{
		Team:      *team,
		GitlabUrl: fmt.Sprintf("/api/v1/classrooms/joined/%s/teams/%s/gitlab", team.ClassroomID.String(), team.ID.String()),
	}

	return c.JSON(response)
}
