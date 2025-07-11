package main

import "fmt"

func numIslands(grid [][]byte) int {
	totalIlhas := 0

	for linha := 0; linha < len(grid); linha++ {
		for coluna := 0; coluna < len(grid[0]); coluna++ {
			if grid[linha][coluna] == '1' {
				buscaProfundidade(grid, linha, coluna)
				totalIlhas++
			}
		}
	}

	return totalIlhas
}

func buscaProfundidade(grid [][]byte, linha, coluna int) {
	if linha < 0 || coluna < 0 || linha >= len(grid) || coluna >= len(grid[0]) || grid[linha][coluna] == '0' {
		return
	}

	grid[linha][coluna] = '0'

	buscaProfundidade(grid, linha-1, coluna) // cima
	buscaProfundidade(grid, linha+1, coluna) // baixo
	buscaProfundidade(grid, linha, coluna-1) // esquerda
	buscaProfundidade(grid, linha, coluna+1) // direita
}

func main() {

	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))
}
