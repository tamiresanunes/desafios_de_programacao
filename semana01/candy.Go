package main

import "fmt"

// Complexidade de tempo e espaço -> O(n)
// O(n) porque o algoritmo percorre a lista de ratings 3 vezes, e cada iteração é linear.

func candy(ratings []int) int {

	// Cria um slice de inteiros chamado "candies" com o mesmo tamanho de "ratings"
	// Esse slice armazenará a quantidade de doces que cada criança receberá
	candies := make([]int, len(ratings))

	// Inicializa o array de doces atribuindo pelo menos 1 doce para cada criança
	for i := range candies {
		candies[i] = 1
	}

	// Primeiro loop: percorre da esquerda para a direita
	// Garante que crianças com nota maior que a anterior recebam mais doces
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] { // Se a nota da criança atual for maior que a anterior
			candies[i] = candies[i-1] + 1 // Atribui 1 doce a mais do que a anterior
		}
		fmt.Println(candies) // Exibe o estado atual do array de doces
	}

	// Segundo loop: percorre da direita para a esquerda
	// Garante que crianças com nota maior que a próxima recebam mais doces
	for i := len(ratings) - 2; i >= 0; i-- {
		// Se a nota da criança for maior que a próxima e ela tem menos ou igual doces que a próxima
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1 // Atribui 1 doce a mais do que a próxima criança
		}
		fmt.Println(candies) // Exibe o estado atualizado do array de doces
	}

	fmt.Println("")       // Apenas para separar a saída no console
	fmt.Println(candies) // Exibe o array final de doces

	// Calcula o total de doces somando todos os valores do array "candies"
	totalCandies := 0
	for _, value := range candies {
		totalCandies += value
	}

	return totalCandies // Retorna o total de doces necessários
}

func main() {
	// Define as notas das crianças
	children := []int{1, 2, 87, 87, 87, 2, 1}

	// Chama a função candy passando as notas como entrada
	resultado := candy(children)

	// Exibe o total de doces necessários
	fmt.Println(resultado)
}
