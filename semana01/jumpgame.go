package main  // Define que este é o pacote principal do programa Go

import "fmt"  // Importa o pacote fmt para entrada e saída de dados

// Função que verifica se é possível alcançar o último índice do array
func canJump(nums []int) bool {

	// jumpCount armazena a posição mais distante que podemos alcançar
	jumpCount := 0

	// Percorre o array verificando até onde é possível chegar
	for i := 0; i < len(nums); i++ {
		// Se a posição atual for maior do que a maior posição alcançável, significa que não é possível continuar
		if i > jumpCount {
			return false  // Retorna false se um índice for inalcançável
		}

		// Atualiza a posição máxima que podemos alcançar a partir do índice atual
		if jumpCount < i+nums[i] {
			jumpCount = i + nums[i]
		}
	}
	return true  // Se percorremos toda a lista, significa que é possível alcançar o final
}

func main() {
	// Teste 1: Deve retornar true (é possível alcançar o final)
	case1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(case1)) // true

	// Teste 2: Deve retornar false (não é possível passar do índice 3)
	case2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(case2)) // false

	// Teste 3: Deve retornar true (apenas um elemento, já estamos no final)
	case3 := []int{0}
	fmt.Println(canJump(case3)) // true

	// Teste 4: Deve retornar true (o primeiro número permite avançar até o final)
	case4 := []int{2, 0, 0}
	fmt.Println(canJump(case4)) // true
}

