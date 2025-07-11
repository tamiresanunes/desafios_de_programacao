package main

import "fmt"

func saoVizinhas(palavra1, palavra2 string) bool {
	diferencas := 0
	for i := 0; i < len(palavra1); i++ {
		if palavra1[i] != palavra2[i] {
			diferencas++
		}
	}
	return diferencas == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visitado := make([]bool, len(wordList))
	fila := []string{beginWord}
	passos := 1

	for len(fila) > 0 {
		tamanhoNivel := len(fila)

		for i := 0; i < tamanhoNivel; i++ {
			palavraAtual := fila[0]
			fila = fila[1:]

			if palavraAtual == endWord {
				return passos
			}

			for j := 0; j < len(wordList); j++ {
				if !visitado[j] && saoVizinhas(palavraAtual, wordList[j]) {
					fila = append(fila, wordList[j])
					visitado[j] = true
				}
			}
		}
		passos++
	}

	return 0
}

func main() {
	inicio := "hit"
	fim := "cog"
	lista := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(ladderLength(inicio, fim, lista))
}
