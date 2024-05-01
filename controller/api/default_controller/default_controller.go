package default_controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gitlab.hs-flensburg.de/gitlab-classroom/model/database/query"
	mailRepo "gitlab.hs-flensburg.de/gitlab-classroom/repository/mail"
	"gitlab.hs-flensburg.de/gitlab-classroom/wrapper/context"
)

type Params struct {
	ClassroomID         *uuid.UUID `params:"classroomId"`
	AssignmentID        *uuid.UUID `params:"assignmentId"`
	AssignmentProjectID *uuid.UUID `params:"projectId"`
	MemberID            *int       `params:"memberId"`
	TeamID              *uuid.UUID `params:"teamId"`
}

type DefaultController struct {
	mailRepo mailRepo.Repository
}

func NewApiController(mailRepo mailRepo.Repository) *DefaultController {
	return &DefaultController{mailRepo: mailRepo}
}

func (ctrl *DefaultController) RotateAccessToken(c *fiber.Ctx, classroom *database.Classroom) error {
	repo := context.Get(c).GetGitlabRepository()
	expiresAt := time.Now().AddDate(0, 0, 364)
	accessToken, err := repo.RotateGroupAccessToken(classroom.GroupID, classroom.GroupAccessTokenID, expiresAt)
	if err != nil {
		return err
	}

	classroom.GroupAccessTokenID = accessToken.ID
	classroom.GroupAccessToken = accessToken.Token
	err = query.Classroom.WithContext(c.Context()).Save(classroom)
	if err != nil {
		return err
	}
	return nil
}
