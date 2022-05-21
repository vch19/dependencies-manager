# Dependency Manager

This dependency manager parses passed **docker-compose** file in **YML** extension, sorts your services topographically
and returns a list of sorted container names or container name by specific index.

## Installation
First, use ```go get``` to install the latest version of the library. 
This command will install the ```dependency-manager``` with the library and its dependencies:

```sh 
~$ go get github.com/vch19/dependencies-manager
```

## Example

```sh
import (
	manager "github.com/vch19/dependencies-manager"
)

func main() {
	//parse yml file to graph
	dockerCompose := manager.ParseYML("PATH_TO_DOCKER_COMPOSE_FILE")

	//sorts the graph topographically and returns 
	//it as an array with the name of these services
	sortedGraph := manager.Sort(&dockerCompose)
}
```
