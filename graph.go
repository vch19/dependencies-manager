package manager

import (
	"log"
)

type Graph struct {
	Vertexes int
	Nodes    map[string][]string
}

func initGraph(vertexes int) Graph {
	return Graph{
		Vertexes: vertexes,
		Nodes:    make(map[string][]string),
	}
}

func (graph *Graph) addVertex(serviceName string) bool {
	if _, ok := graph.Nodes[serviceName]; ok {
		log.Printf("serviceName %v already exists\n", serviceName)
		return false
	}

	graph.Nodes[serviceName] = []string{}
	return true
}

func (graph *Graph) addEdge(serviceName string, dependsOn string) bool {
	if _, ok := graph.Nodes[serviceName]; !ok {
		graph.addVertex(dependsOn)
	}

	graph.Nodes[serviceName] = append(graph.Nodes[serviceName], dependsOn)
	return true
}

func (dockerCompose *DockerCompose) Sort() []string {
	graph := initGraph(len(dockerCompose.Services))

	for key, value := range dockerCompose.Services {
		graph.addVertex(key)
		for _, dependency := range value.DependsOn {
			graph.addEdge(key, dependency)
		}
	}

	return graph.topologicalSort()
}

func (graph *Graph) topologicalSort() []string {
	var stack []string
	var sortedGraph []string

	visited := make(map[string]bool)

	for key := range graph.Nodes {
		visited[key] = false
	}

	for key := range graph.Nodes {
		if !visited[key] {
			topologicalSortUtil(key, visited, &stack, graph)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		sortedGraph = append(sortedGraph, stack[i])
	}

	return sortedGraph
}

func topologicalSortUtil(vertex string, visited map[string]bool, stack *[]string, graph *Graph) {
	visited[vertex] = true
	childNodes := graph.Nodes[vertex]

	for i := range childNodes {
		if len(childNodes) > 0 && !visited[childNodes[i]] {
			topologicalSortUtil(childNodes[i], visited, stack, graph)
		}
	}

	*stack = append(*stack, vertex)
}
