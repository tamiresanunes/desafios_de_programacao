package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Criar o grafo como lista de adjacência
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		a, b := pair[0], pair[1]
		graph[b] = append(graph[b], a)
	}

	// Estados de visitação:
	// 0 = não visitado
	// 1 = visitando (pilha de recursão)
	// 2 = visitado
	visited := make([]int, numCourses)

	// DFS
	var dfs func(int) bool
	dfs = func(course int) bool {
		if visited[course] == 1 {
			// ciclo detectado
			return false
		}
		if visited[course] == 2 {
			// já processado
			return true
		}
		// marca como "visitando"
		visited[course] = 1
		for _, next := range graph[course] {
			if !dfs(next) {
				return false
			}
		}
		// marca como "visitado"
		visited[course] = 2
		return true
	}

	// Verifica todos os cursos
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

// Teste
func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))          // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))  // false
}
