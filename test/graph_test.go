package test

import (
	manager "github.com/richard2259/dependencies-manager"
	"os"
	"testing"
)

func TestSortGraph(t *testing.T) {
	dockerCompose := manager.ParseYML(os.Getenv("TEST_DOCKER_COMPOSE_FILE_PATH"))
	sortedGraph := dockerCompose.Sort()

	if len(sortedGraph) != len(dockerCompose.Services) && len(sortedGraph) != 0 {
		t.Errorf("The length of the sorted graph must be equal to the docker-compose.Services and both of them mustn't equal 0")
	}
}