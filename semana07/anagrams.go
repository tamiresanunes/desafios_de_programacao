package main

import (
	"fmt"
	"sort"
)

/*
Complexidade de tempo:
sortString() custa O(k log k), onde k é o tamanho da palavra
Laço externo e interno: até O(n²) iterações
Tempo total = O(n² × k log k)

Complexidade de espaço:
Espaço total = O(n × k)


Forma mais inteligente de resolver: usando map
*/

func groupAnagrams(strs []string) [][]string {
	agrupados := [][]string{}
	visitado := make([]bool, len(strs))

	sortString := func(s string) string {
		r := []rune(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		return string(r)
	}

	for i := 0; i < len(strs); i++ {
		if visitado[i] {
			continue
		}

		palavraBase := strs[i]
		palavraBaseOrdenada := sortString(palavraBase)
		grupo := []string{palavraBase}
		visitado[i] = true

		for j := i + 1; j < len(strs); j++ {
			if visitado[j] {
				continue
			}

			if sortString(strs[j]) == palavraBaseOrdenada {
				grupo = append(grupo, strs[j])
				visitado[j] = true
			}
		}

		agrupados = append(agrupados, grupo)
	}

	return agrupados

}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
