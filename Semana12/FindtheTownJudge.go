func findJudge(n int, trust [][]int) int {
    // Caso especial: se não houver relações de confiança e houver apenas uma pessoa,
    // ela é o juiz por padrão
    if len(trust) == 0 && n == 1 {
        return 1
    }

    // Inicializa um array de pontuação de confiança
    trustScore := make([]int, n+1) // índice 0 é ignorado

    // Processa cada relação de confiança
    for _, t := range trust {
        a, b := t[0], t[1]
        trustScore[a]--  // a confia em alguém → não é o juiz
        trustScore[b]++  // b é confiado por alguém
    }

    // Procura por alguém com pontuação n - 1
    for i := 1; i <= n; i++ {
        if trustScore[i] == n-1 {
            return i
        }
    }

    return -1 // Nenhum juiz encontrado
}
