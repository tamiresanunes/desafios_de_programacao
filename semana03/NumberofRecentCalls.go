package main

import (
	"fmt"
)

// RecentCounter struct
type RecentCounter struct {
	requests []int
}

// Construtor
func Constructor() RecentCounter {
	return RecentCounter{requests: []int{}}
}

// Método ping
func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	// Remove todos os tempos menores que t - 3000
	for this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	// Retorna o número de elementos dentro da janela [t-3000, t]
	return len(this.requests)
}

// Teste da implementação
func main() {
	recentCounter := Constructor()
	fmt.Println(recentCounter.Ping(1))    // return 1
	fmt.Println(recentCounter.Ping(100))  // return 2
	fmt.Println(recentCounter.Ping(3001)) // return 3
	fmt.Println(recentCounter.Ping(3002)) // return 3
}
