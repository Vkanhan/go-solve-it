package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	vertices map[string][]string
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{vertices: make(map[string][]string)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 string) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1) // For undirected graph
}

// DFS performs a depth-first search
func (g *Graph) DFS(start string, visited map[string]bool) {
	visited[start] = true
	fmt.Println(start) // Process the node

	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("B", "E")
	graph.AddEdge("C", "F")
	graph.AddEdge("E", "F")

	fmt.Println("DFS starting from node A:")
	visited := make(map[string]bool)
	graph.DFS("A", visited)
}
