// Breadth-First Search

package main

import (
	"container/list"
	"fmt"
)

func main() {
	graph := NewGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)

	startVertex := 0
	bfsResult := BFS(graph, startVertex)

	fmt.Printf("BFS comenzando desde el nodo %d: %v\n", startVertex, bfsResult)
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

func BFS(graph *Graph, start int) []int {
	visited := make(map[int]bool)
	result := make([]int, 0)

	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(int)
		result = append(result, current)

		for _, neighbor := range graph.AdjList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}

	return result
}
