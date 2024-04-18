package utils

import (
	"log"

	dbModel "gitlab.hs-flensburg.de/gitlab-classroom/model/database"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	log.Println("Running database migrations")
	return db.AutoMigrate(
		&dbModel.User{},
		&dbModel.UserAvatar{},
		&dbModel.Classroom{},
		&dbModel.Team{},
		&dbModel.UserClassrooms{},
		&dbModel.Assignment{},
		&dbModel.AssignmentProjects{},
		&dbModel.ClassroomInvitation{},
	)
}
