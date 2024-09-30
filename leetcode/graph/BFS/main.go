package main

import (
	"container/list"
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

// BFS performs a breadth-first search
func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(string)
		fmt.Println(node) // Process the node

		for _, neighbor := range g.vertices[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
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

	fmt.Println("BFS starting from node A:")
	graph.BFS("A")
}
