package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {

	mapa := make(map[int][]int)
	for _, par := range prerequisites {
		curso := par[0]
		pre := par[1]
		mapa[curso] = append(mapa[curso], pre)
	}
	estado := make(map[int]string)

	for curso := 0; curso < numCourses; curso++ {
		if !dfs(curso, mapa, estado) {
			fmt.Println("Não é possível concluir todos os cursos (ciclo detectado).")
			return false
		}
	}

	fmt.Println("Todos os cursos podem ser concluídos!")
	return true
}

func dfs(curso int, mapa map[int][]int, estado map[int]string) bool {

	fmt.Printf("\nAnalisando o curso %d\n", curso)
	if estado[curso] == "visitando" {
		fmt.Printf("Curso %d já está sendo visitado. Ciclo detectado!\n", curso)
		return false
	}
	if estado[curso] == "concluído" {
		fmt.Printf("Curso %d já foi concluído. Ignorando.\n", curso)
		return true
	}

	fmt.Printf("Marcando curso %d como 'visitando'\n", curso)
	estado[curso] = "visitando"

	for _, prerequisito := range mapa[curso] {
		fmt.Printf("Curso %d depende de %d\n", curso, prerequisito)
		if !dfs(prerequisito, mapa, estado) {
			return false
		}
	}

	estado[curso] = "concluído"
	return true
}

func main() {
	prerequisites := [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
	}

	numCourses := 4

	fmt.Println(canFinish(numCourses, prerequisites))
}
