package manager_test

import (
	manager "github.com/richard2259/dependencies-manager"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseYMLToDockerCompose(t *testing.T) {
	dockerCompose := manager.ParseYML("data/docker-compose.yml")

	require.FileExists(t, "data/docker-compose.yml")
	require.NotEmpty(t, dockerCompose.Version)
	require.Len(t, dockerCompose.Services["chrome-browser"].DependsOn, 2)
}
