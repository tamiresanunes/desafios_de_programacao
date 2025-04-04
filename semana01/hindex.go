package main  // Define que este é o pacote principal do programa Go

import (
	"fmt"   // Pacote usado para entrada e saída de dados (print no console)
	"sort"  // Pacote usado para ordenação de slices
)

// Função que calcula o índice H de um pesquisador com base na lista de citações
func hIndex(citations []int) int {

	sort.Ints(citations)  // Ordena o slice de citações em ordem crescente

	hIndex := 0   // Inicializa o índice H como 0
	articles := 1 // Contador de artigos processados, começa com 1

	// Percorre a lista de trás para frente (do maior para o menor)
	for i := len(citations) - 1; i >= 0; i-- {
		// Verifica se o número de artigos é menor ou igual ao número de citações
		if articles <= citations[i] {
			articles++  // Aumenta a contagem de artigos considerados
			hIndex++    // Atualiza o índice H
		} else {
			break  // Se a condição falhar, interrompe o loop
		}
	}
	return hIndex // Retorna o índice H calculado
}

// Função principal do programa
func main() {

	citations := []int{3, 0, 6, 1, 5}  // Declara um slice com as citações de artigos
	fmt.Println(hIndex(citations))     // Calcula e imprime o índice H no console
}

