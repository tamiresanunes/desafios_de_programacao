package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	clonados := make(map[*Node]*Node) //chave: nó original; valor: nó clonado
	return clonarNo(node, clonados)
}

func clonarNo(no *Node, clonados map[*Node]*Node) *Node {
	if existente, ok := clonados[no]; ok {
		return existente
	}

	copia := &Node{Val: no.Val}
	clonados[no] = copia

	for _, vizinho := range no.Neighbors {
		copia.Neighbors = append(copia.Neighbors, clonarNo(vizinho, clonados))
	}

	return copia
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	clone := cloneGraph(n1)

	fmt.Println("Clone:", clone.Val)
	for _, vizinho := range clone.Neighbors {
		fmt.Println(" Vizinhos: ", vizinho.Val)
	}
}
