package main

import "fmt"

func isValid(s string) bool {

	var pilha []string
	var isValid bool

	for i := 0; i < len(s); i++ {

		isValid = false
		if len(s) == 0 || len(s) == 1 || len(s)%2 != 0 {
			return false
		}

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {

			pilha = append(pilha, string(s[i]))

		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {

			if len(pilha) == 0 {
				return false
			}

			topo := pilha[len(pilha)-1]

			if (s[i] == ')' && topo != "(") || (s[i] == ']' && topo != "[") ||
				(s[i] == '}' && topo != "{") {

				return false
			}
			isValid = true
			pilha = pilha[:len(pilha)-1]
		}

	}
	if len(pilha) != 0 {
		return false
	}

	return isValid
}

func main() {
	s := "[[[]"
	fmt.Println(isValid(s))
}

// Complexidade de tempo: O(n)
// Complexidade de espaço: O(n)
// Existe soluçãi mais eficiente: Não, o problema exige analisar
// 								  todos os caracteres da string
//                                pelo menos uma vez
// Estruturas utilizadas: Pilha
