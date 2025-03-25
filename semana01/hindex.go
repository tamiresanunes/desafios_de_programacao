package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {

	sort.Ints(citations)
	hIndex := 0
	articles := 1

	for i := len(citations) - 1; i >= 0; i-- {
		if articles <= citations[i] {
			articles++
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

func main() {

	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
