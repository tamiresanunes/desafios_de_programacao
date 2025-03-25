package main

import (
	"fmt"
	"strconv"
)

// Complexidade de Tempo:  O(n) -> n = 32
// Complexidade de Espaço: O(n) -> n = 32

func reverseBits(num uint32) uint32 {

	binaryString := fmt.Sprintf("%032b", num) 
	fmt.Println(binaryString)
	reverseSlice := make([]byte, 32)

	for i := 0; i < 32; i++ {
		reverseSlice[i] = binaryString[31-i]
	}

	reverseStr := string(reverseSlice)
	fmt.Println(reverseStr)
	result, _ := strconv.ParseUint(reverseStr, 2, 32)
	return uint32(result)
}

func main() {
	output := reverseBits(43261596)
	fmt.Println("Resultado ", output)
}