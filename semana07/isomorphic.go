package main

import "fmt"

/*
Complexidade de tempo = O(n)
Complexidade de espaço = O(n)
*/

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	mapST := map[rune]rune{}
	mapTS := map[rune]rune{}

	for i := range s {
		letraT := rune(t[i])
		letraS := rune(s[i])

		if letra, existe := mapST[letraS]; existe {
			if letra != letraT {
				return false
			}
		} else {
			mapST[letraS] = letraT
		}

		if letra, existe := mapTS[letraT]; existe {
			if letra != letraS {
				return false
			}
		} else {
			mapTS[letraT] = letraS
		}
	}

	return true
}

func main() {

	s := "egg"
	t := "add"

	fmt.Println(isIsomorphic(s, t))

}
