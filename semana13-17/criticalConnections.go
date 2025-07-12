package main

import (
	"fmt"
)

func criticalConnections(n int, connections [][]int) [][]int {
	// Grafo como lista de adjacência
	graph := make([][]int, n)
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Inicializações
	disc := make([]int, n)
	low := make([]int, n)
	time := 1
	var result [][]int

	// DFS
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		disc[u] = time
		low[u] = time
		time++

		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			if disc[v] == 0 {
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if low[v] > disc[u] {
					result = append(result, []int{u, v})
				}
			} else {
				low[u] = min(low[u], disc[v])
			}
		}
	}

	dfs(0, -1)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Exemplo de uso
func main() {
	n1 := 4
	connections1 := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}
	fmt.Println("Output 1:", criticalConnections(n1, connections1)) // [[1 3]]

	n2 := 2
	connections2 := [][]int{{0, 1}}
	fmt.Println("Output 2:", criticalConnections(n2, connections2)) // [[0 1]]
}
