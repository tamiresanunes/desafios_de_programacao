package main

import "fmt"

//Complexidade de tempo: O(n) - São percorridos todos os nós da árvore
//Complexidade de espaço: Arvore perfeitamente balanceada -> O(h) h = altura da arvore
//                       Arvore totalmente desbalancedada -> O(n) n = numero de nós
// Dá para fazer de forma mais eficiente quanto ao espaço utilizando uma abordagem
// iterativa como Busca em Largura
// Na abordagem iterativa pode ser utilizada fila

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {

	root := &TreeNode{
		val: 3,
		Left: &TreeNode{
			val: 9,
		},
		Right: &TreeNode{
			val: 20,
			Left: &TreeNode{
				val: 15,
			},
			Right: &TreeNode{
				val: 7,
			},
		},
	}

	fmt.Println(maxDepth(root))

}
