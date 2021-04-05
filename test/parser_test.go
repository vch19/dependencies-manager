package test

import (
	manager "github.com/richard2259/dependencies-manager"
	"os"
	"testing"
)

const numberOfServices int = 12

func TestParseYMLToDockerCompose(t *testing.T) {
	dockerCompose := manager.ParseYML(os.Getenv("TEST_DOCKER_COMPOSE_FILE_PATH"))
	if dockerCompose.Version == "" {
		t.Error("Version is empty!")
	}
	if len(dockerCompose.Services) != numberOfServices {
		t.Errorf("Current number of services %b.\nShould be:%b", len(dockerCompose.Services), numberOfServices)
	}
	if len(dockerCompose.Services["database"].DependsOn) == 0 {
		t.Errorf("The 'database' service shoud depends on: %b services, but currently on: %b",
			1, len(dockerCompose.Services["database"].DependsOn))
	}
}
