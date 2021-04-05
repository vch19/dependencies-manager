package manager

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
		log.Fatal("Error occurred during file reading")
	}

	services := DockerCompose{}

	parsingErr := yaml.Unmarshal(bytes, &services)

	if parsingErr != nil {
		log.Fatalf("Error occurred during unmarshaling yml file")
	}

	return services
}
