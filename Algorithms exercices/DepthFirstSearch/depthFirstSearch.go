// Depth-First Search

package main

import (
	"fmt"
)

func main() {
	graph := NewGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)

	startVertex := 0
	DFSTraversal(graph, startVertex)
}

type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

func DFS(graph *Graph, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range graph.AdjList[start] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}

func DFSTraversal(graph *Graph, start int) {
	visited := make(map[int]bool)
	fmt.Print("DFS comenzando desde el nodo ", start, ": ")
	DFS(graph, start, visited)
	fmt.Println()
}
