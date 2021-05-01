package manager

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DockerCompose struct {
	Version  string
	Services map[string]Service
}

type Service struct {
	DependsOn []string `yaml:"depends_on"`
}

func ParseYML(path string) DockerCompose {
	bytes, ioErr := ioutil.ReadFile(path)

	if ioErr != nil {
		fmt.Println("Error occurred during file reading")
		return DockerCompose{}
	}

	var dockerCompose DockerCompose

	parsingErr := yaml.Unmarshal(bytes, &dockerCompose)

	if parsingErr != nil {
		fmt.Println("Error occurred during unmarshalling YML file")
		return dockerCompose
	}

	return dockerCompose
}
