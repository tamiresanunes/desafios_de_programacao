package main

import "fmt"

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return false // Cycle detected
	}
	uf.parent[rootX] = rootY
	return true
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n)

	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}

func main() {
	edges1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	edges2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}

	fmt.Println("Output 1:", findRedundantConnection(edges1)) // [2 3]
	fmt.Println("Output 2:", findRedundantConnection(edges2)) // [1 4]
}
