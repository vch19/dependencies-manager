package manager_test

import (
	"fmt"
	manager "github.com/richard2259/dependencies-manager"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortGraph(t *testing.T) {
	dockerCompose := manager.ParseYML("data/docker-compose.yml")
	firstTimeSorted := manager.Sort(&dockerCompose)
	secondTimeSorted := manager.Sort(&dockerCompose)

	fmt.Println(firstTimeSorted)
	fmt.Println(secondTimeSorted)

	require.FileExists(t, "data/docker-compose.yml")
	require.Equal(t, firstTimeSorted, secondTimeSorted)
}
