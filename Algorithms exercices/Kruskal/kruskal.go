// Kruskal's algorithm

package main

import (
	"fmt"
	"sort"
)

func main() {
	graph := Graph{
		Vertices: 4,
		Edges:    5,
		EdgeList: []Edge{
			{0, 1, 10},
			{0, 2, 6},
			{0, 3, 5},
			{1, 3, 15},
			{2, 3, 4},
		},
	}

	minimumSpanningTree := kruskal(&graph)

	fmt.Println("Árbol de expansión mínimo:")
	for _, edge := range minimumSpanningTree {
		fmt.Printf("Arista: %d - %d, Peso: %d\n", edge.Start, edge.End, edge.Weight)
	}
}

type Edge struct {
	Start, End, Weight int
}

type Graph struct {
	Vertices, Edges int
	EdgeList        []Edge
}

type UnionFind struct {
	parent []int
	rank   []int
}

func newUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{parent, rank}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func kruskal(graph *Graph) []Edge {
	sort.Slice(graph.EdgeList, func(i, j int) bool {
		return graph.EdgeList[i].Weight < graph.EdgeList[j].Weight
	})

	mst := make([]Edge, 0)
	uf := newUnionFind(graph.Vertices)

	for _, edge := range graph.EdgeList {
		if uf.find(edge.Start) != uf.find(edge.End) {
			mst = append(mst, edge)
			uf.union(edge.Start, edge.End)
		}
	}

	return mst
}
