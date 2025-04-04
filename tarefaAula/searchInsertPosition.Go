func searchInsert(nums []int, target int) int {
    //Criamos dois ponteiros: início e fim, respectivamente, início e fim do array
    //usamos len para adptar o tamanho do array, para pegarmos o índice certo
    left, right := 0, len(nums)-1

    //O laço vai ser executado enquanto houver elemntos para comprar no intervalo válido
    for left <= right {
        // Isso vai calcular o índice do meio do intervalo atual
        mid := left + (right-left)/2

        //Aqui vamos verificar se o valor do meio é exatamente o alvo
        if nums[mid] == target {
            return mid //Retorna o índice se encontrar
        } else if nums[mid] < target {
            // O alvo está na metade direita, então ignoramos a esquerda
            left = mid + 1
        } else {
            // O alvo está na metade esquerda, então ignoramos a direita
            right = mid - 1
        }
    }

    // Se sair do laço, não encontrou o target.
    // Nesse ponto, 'left' será a posição correta onde o valor deveria ser inserido
    return left
}