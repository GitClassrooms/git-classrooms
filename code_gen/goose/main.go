package main

import (
	"context"

	"github.com/pressly/goose/v3"
	"gitlab.hs-flensburg.de/gitlab-classroom/config"
	_ "gorm.io/driver/postgres"

	"log"
	"os"
)

const (
	mode int = iota + 1
	command
	minArgLength
)

var modeDir = map[string]string{
	"normal": "model/database/migrations",
	"seed":   "model/database/seeds",
}

func main() {
	args := os.Args

	cfg, err := config.LoadApplicationConfig()
	if err != nil {
		log.Fatalf("failed to load application config: %v", err)
	}

	if len(args) < minArgLength {
		log.Fatalf("goose: invalid number of arguments")
	}

	mode := args[mode]
	command := args[command]

	dir, ok := modeDir[mode]
	if !ok {
		log.Fatalf("goose: invalid mode")
	}

	db, err := goose.OpenDBWithDriver("postgres", cfg.Database.Dsn())
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > minArgLength {
		arguments = append(arguments, args[minArgLength:]...)
	}

	if err := goose.RunContext(context.Background(), command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
