package main

import (
	"container/list"
	"fmt"
)

//Complexidade de tempo: O(n) n= numero de nós da árvore
//Complexidade de espaço: O(n)
// Não há forma mais eficiente porque devemos visitar todos os nós da árvore
// pelo mens uma vez.
// Estruturas utilizadas: fila e busca em largura (BFS)
// Estruturas que podem wser utilizadas: pilha e busca em profundidade (DFS)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		level := []int{}

		for i := 0; i < levelSize; i++ {
			element := queue.Front()
			queue.Remove(element)
			node := element.Value.(*TreeNode)

			level = append(level, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, level)
	}
	return result
}

func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrder(root))
}
