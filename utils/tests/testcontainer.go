package tests

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartPostgres() (*postgres.PostgresContainer, error) {
	return postgres.Run(context.Background(),
		"postgres:14-alpine",
		postgres.WithDatabase("classrooms_db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}

func GetPostgresData(postgres testcontainers.Container, err error) (testcontainers.Container, string, error) {
	if err != nil {
		return nil, "", err
	}

	postgresHost, err := postgres.ContainerIP(context.Background())
	if err != nil {
		return nil, "", err
	}

	return postgres, postgresHost, nil
}
