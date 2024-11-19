package test_utils

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPgContainer() (string, string, func(), error) {
	ctx := context.Background()

	// Configure the PostgreSQL test container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:17",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return "", "", nil, err
	}

	// Get the mapped host and port
	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return "", "", nil, err
	}
	hostIP, err := postgresContainer.Host(ctx)
	if err != nil {
		return "", "", nil, err
	}

	cleanup := func() {
		postgresContainer.Terminate(ctx)
	}

	return hostIP, mappedPort.Port(), cleanup, nil
}
