package main

import (
    "context"
    "testing"
    "time"

    "github.com/ory/dockertest/v3"
    "github.com/ory/dockertest/v3/docker"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetMongoClient(t *testing.T) {
    // Use the dockertest library to start a MongoDB instance
    pool, err := dockertest.NewPool("")
    if err != nil {
        t.Fatalf("Could not connect to Docker: %s", err)
    }

    resource, err := pool.RunWithOptions(&dockertest.RunOptions{
        Repository: "mongo",
        Tag:        "latest",
        PortBindings: map[docker.Port][]docker.PortBinding{
            "27017/tcp": {{HostIP: "0.0.0.0", HostPort: "27017"}},
        },
    })
    if err != nil {
        t.Fatalf("Could not start resource: %s", err)
    }

    // The application in the container might not be ready to accept connections yet,
    // so we retry a few times
    if err := pool.Retry(func() error {
        var err error
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        _, err = getMongoClient("mongodb://localhost:27017", ctx)
        if err != nil {
            return err
        }

        return nil
    }); err != nil {
        t.Fatalf("Could not connect to Docker: %s", err)
    }

    // When you're done, kill and remove the container
    if err := pool.Purge(resource); err != nil {
        t.Fatalf("Could not purge resource: %s", err)
    }
}
