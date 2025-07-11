package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	linhas := len(heights)
	colunas := len(heights[0])

	// matriz visitado no dfs
	pacifico := make([][]bool, linhas)
	atlantico := make([][]bool, linhas)
	for i := 0; i < linhas; i++ {
		pacifico[i] = make([]bool, colunas)
		atlantico[i] = make([]bool, colunas)
	}
	fmt.Println("Matriz pacifico:")
	for i := 0; i < linhas; i++ {
		fmt.Println(pacifico[i])
	}

	for linha := 0; linha < linhas; linha++ {
		dfs(linha, 0, pacifico, heights[linha][0], heights)
		dfs(linha, colunas-1, atlantico, heights[linha][colunas-1], heights)
	}

	for coluna := 0; coluna < colunas; coluna++ {
		dfs(0, coluna, pacifico, heights[0][coluna], heights)
		dfs(linhas-1, coluna, atlantico, heights[linhas-1][coluna], heights)
	}

	resultado := [][]int{}
	for linha := 0; linha < linhas; linha++ {
		for coluna := 0; coluna < colunas; coluna++ {
			if pacifico[linha][coluna] && atlantico[linha][coluna] {
				resultado = append(resultado, []int{linha, coluna})
			}
		}
	}

	return resultado
}

func dfs(linha, coluna int, visitado [][]bool, alturaAnterior int, heights [][]int) {
	linhas := len(heights)
	colunas := len(heights[0])

	if linha < 0 || coluna < 0 || linha >= linhas || coluna >= colunas {
		return
	}

	if visitado[linha][coluna] || heights[linha][coluna] < alturaAnterior {
		return
	}

	visitado[linha][coluna] = true

	dfs(linha+1, coluna, visitado, heights[linha][coluna], heights)
	dfs(linha-1, coluna, visitado, heights[linha][coluna], heights)
	dfs(linha, coluna+1, visitado, heights[linha][coluna], heights)
	dfs(linha, coluna-1, visitado, heights[linha][coluna], heights)
}

func main() {
	matriz := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	fmt.Println(pacificAtlantic(matriz))
}
