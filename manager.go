package manager

import "log"

func ReturnAllSortedServices(path string) []string {
	dockerCompose := ParseYML(path)
	return dockerCompose.Sort()
}

func Return(path string, index int) string {
	dockerCompose := ParseYML(path)
	sortedGraph := dockerCompose.Sort()

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
