package router

import (
	"de.hs-flensburg.gitlab/gitlab-classroom/config"
	"de.hs-flensburg.gitlab/gitlab-classroom/handler"
	apiHandler "de.hs-flensburg.gitlab/gitlab-classroom/handler/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App, applicationConfig *config.Config) {
	app.Static("/", applicationConfig.FrontendPath)

	app.Get("/auth", handler.Auth)
	app.Get(applicationConfig.Auth.RedirectURL.Path, handler.Callback)

	api := app.Group("/api", logger.New()) // behind "/api" is always a user logged into the session and this user is logged into the repository, which is accessable via "ctx.Locals("gitlab-repo").(repository.Repository)"

	fiberHandler := apiHandler.NewFiberApiHandler()

	api.Post("/createClassroom", fiberHandler.CreateClassroom)
	api.Post("/createAssignment", fiberHandler.CreateAssignment)
}
