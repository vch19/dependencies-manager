package manager

func ReturnAllSortedServices(path string) []string {
	dockerCompose := ParseYML(path)
	return Sort(&dockerCompose)
}
