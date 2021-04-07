package cmd

import (
	"github.com/richard2259/dependencies-manager"
	"log"
)

func ReturnAllSortedServices(path string) []string {
	dockerCompose := manager.ParseYML(path)
	return manager.Sort(&dockerCompose)
}

func Return(path string, index int) string {
	dockerCompose := manager.ParseYML(path)
	sortedGraph := manager.Sort(&dockerCompose)

	return get(sortedGraph, index)
}

func get(slice []string, index int) string {
	for i := 0; i < len(slice); i++ {
		if i == index {
			return slice[i]
		}
	}
	log.Println("Index doesn't exist in the passed array")
	return ""
}
