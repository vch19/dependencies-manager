package manager

import (
	"fmt"
)

func ReturnAllSortedServices(path string) []string {
	dockerCompose := ParseYML(path)
	return Sort(&dockerCompose)
}

func Return(path string, index int) string {
	dockerCompose := ParseYML(path)
	sortedGraph := Sort(&dockerCompose)

	return get(sortedGraph, index)
}

func get(slice []string, index int) string {
	for i := 0; i < len(slice); i++ {
		if i == index {
			return slice[i]
		}
	}
	fmt.Println("Index doesn't exist in the passed array")
	return ""
}
