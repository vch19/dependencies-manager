package manager

import (
	"fmt"
	"sort"
)

type Graph struct {
	Vertexes int
	Nodes    map[string][]string
	Order    []string
}

func Sort(dockerCompose *DockerCompose) []string {
	graph := initGraph(len(dockerCompose.Services))
	getSortedKeys(dockerCompose.Services, &graph.Order)

	for _, vertex := range graph.Order {
		graph.addVertex(vertex)
		for _, dependency := range dockerCompose.Services[vertex].DependsOn {
			graph.addEdge(vertex, dependency)
		}
	}

	return graph.topologicalSort()
}

func initGraph(vertexes int) Graph {
	return Graph{
		Vertexes: vertexes,
		Nodes:    make(map[string][]string),
	}
}

func (graph *Graph) addVertex(serviceName string) bool {
	if _, ok := graph.Nodes[serviceName]; ok {
		fmt.Printf("serviceName %v already exists\n", serviceName)
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

func getSortedKeys(services map[string]Service, keySet *[]string) {
	for key := range services {
		*keySet = append(*keySet, key)
	}
	sort.Strings(*keySet)
}

func (graph *Graph) topologicalSort() []string {
	var stack []string
	var sortedGraph []string

	visited := make(map[string]bool)

	for key := range graph.Nodes {
		visited[key] = false
	}

	for _, vertex := range graph.Order {
		if !visited[vertex] {
			topologicalSortUtil(vertex, visited, &stack, graph)
		}
	}

	for _, service := range stack {
		sortedGraph = append(sortedGraph, service)
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
