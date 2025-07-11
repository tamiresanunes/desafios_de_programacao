package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	grafo := make(map[int][]int)

	for _, aresta := range edges {
		u := aresta[0]
		v := aresta[1]
		grafo[u] = append(grafo[u], v)
		grafo[v] = append(grafo[v], u)
	}

	folhas := []int{}
	for no, vizinhos := range grafo {
		if len(vizinhos) == 1 {
			folhas = append(folhas, no)
		}
	}

	restantes := n
	for restantes > 2 {
		quantidade := len(folhas)
		restantes -= quantidade
		novasFolhas := []int{}

		for _, folha := range folhas {

			vizinho := grafo[folha][0]

			novos := []int{}
			for _, v := range grafo[vizinho] {
				if v != folha {
					novos = append(novos, v)
				}
			}
			grafo[vizinho] = novos

			if len(grafo[vizinho]) == 1 {
				novasFolhas = append(novasFolhas, vizinho)
			}
		}

		folhas = novasFolhas
	}

	return folhas
}

func main() {
	n := 6
	arestas := [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}

	centros := findMinHeightTrees(n, arestas)
	fmt.Println(centros)
}
