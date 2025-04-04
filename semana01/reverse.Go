package main

import (
	"fmt"
	"strconv"
)

// Complexidade de Tempo:  O(n) -> n = 32
// Complexidade de Espaço: O(n) -> n = 32

func reverseBits(num uint32) uint32 {

	// Converte o número inteiro para uma string binária de 32 bits
	// "%032b" -> Formata o número como binário com 32 dígitos (adicionando zeros à esquerda se necessário)
	binaryString := fmt.Sprintf("%032b", num)  
	fmt.Println(binaryString) // Exibe a representação binária original no console

	// Cria um array de bytes de tamanho 32 (reverseSlice) para armazenar os bits invertidos
	reverseSlice := make([]byte, 32)

	// Percorre os 32 bits e os armazena na ordem inversa no slice
	for i := 0; i < 32; i++ {
		reverseSlice[i] = binaryString[31-i] // Copia os bits da direita para a esquerda
	}

	// Converte o slice de bytes de volta para uma string
	reverseStr := string(reverseSlice)
	fmt.Println(reverseStr) // Exibe a string binária invertida no console

	// Converte a string binária invertida de volta para um número inteiro
	// strconv.ParseUint(string, base 2, tamanho 32 bits)
	result, _ := strconv.ParseUint(reverseStr, 2, 32)

	// Retorna o número invertido convertido para uint32
	return uint32(result)
}

func main() {
	// Define o número de entrada
	output := reverseBits(43261596)

	// Exibe o resultado final
	fmt.Println("Resultado ", output)
}
