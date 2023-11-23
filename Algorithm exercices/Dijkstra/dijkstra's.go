// Dijkstra's algorithm.

package main

import (
	"fmt"
	"math"
)

func main() {
	graph := newGraph()

	graph.addNode("A", map[string]int{"B": 1, "C": 4})
	graph.addNode("B", map[string]int{"A": 1, "C": 2, "D": 5})
	graph.addNode("C", map[string]int{"A": 4, "B": 2, "D": 1})
	graph.addNode("D", map[string]int{"B": 5, "C": 1})

	startNode := "A"
	endNode := "D"

	distance, path, err := dijkstra(graph, startNode, endNode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("La distancia mas corta de %s a %s es %d\n", startNode, endNode, distance)
	fmt.Printf("Camino: %v\n", path)
}

const infinity = math.MaxInt32

type graph struct {
	nodes map[string]map[string]int
}

func newGraph() *graph {
	return &graph{
		nodes: make(map[string]map[string]int),
	}
}

func (g *graph) addNode(name string, neighbors map[string]int) error {
	if _, exists := g.nodes[name]; exists {
		return fmt.Errorf("node %s already exists", name)
	}
	g.nodes[name] = neighbors
	return nil
}

func initializeDistances(nodes []string) map[string]int {
	distances := make(map[string]int)
	for _, node := range nodes {
		distances[node] = infinity
	}
	return distances
}

func updateDistances(current string, graph *graph, distances map[string]int, previous map[string]string) {
	for neighbor, weight := range graph.nodes[current] {
		if alt := distances[current] + weight; alt < distances[neighbor] {
			distances[neighbor] = alt
			previous[neighbor] = current
		}
	}
}

func dijkstra(graph *graph, startNode, endNode string) (int, []string, error) {
	if _, exists := graph.nodes[startNode]; !exists {
		return 0, nil, fmt.Errorf("start node %s not found", startNode)
	}
	if _, exists := graph.nodes[endNode]; !exists {
		return 0, nil, fmt.Errorf("end node %s not found", endNode)
	}

	distances := initializeDistances(nodes(graph))
	previous := make(map[string]string)
	unvisited := make(map[string]struct{})

	for node := range graph.nodes {
		unvisited[node] = struct{}{}
	}

	distances[startNode] = 0

	for len(unvisited) > 0 {
		var current string
		minDist := infinity

		for node := range unvisited {
			if distances[node] < minDist {
				minDist = distances[node]
				current = node
			}
		}

		delete(unvisited, current)
		updateDistances(current, graph, distances, previous)
	}

	path := []string{endNode}
	for previous[endNode] != "" {
		path = append([]string{previous[endNode]}, path...)
		endNode = previous[endNode]
	}

	return distances[endNode], path, nil
}

func nodes(g *graph) []string {
	var result []string
	for node := range g.nodes {
		result = append(result, node)
	}
	return result
}
