package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var path []int

	var dfs func(node int)
	dfs = func(node int) {
		path = append(path, node)

		if node == len(graph)-1 {
			// Chegamos ao destino, adicionamos uma cópia do caminho
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		} else {
			// Continua explorando os vizinhos
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}

		// Backtrack
		path = path[:len(path)-1]
	}

	dfs(0)
	return result
}

// Testando com os exemplos fornecidos
func main() {
	graph1 := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println("Output 1:", allPathsSourceTarget(graph1)) // [[0 1 3] [0 2 3]]

	graph2 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println("Output 2:", allPathsSourceTarget(graph2)) // [[0 4] [0 3 4] [0 1 3 4] [0 1 2 3 4] [0 1 4]]
}
