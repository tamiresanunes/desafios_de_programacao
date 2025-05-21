func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }

    charIndex := make(map[byte]int) // Mapeia caractere para seu índice mais recente
    maxLength := 0
    start := 0

    for end := 0; end < n; end++ {
        if idx, ok := charIndex[s[end]]; ok && idx >= start {
            start = idx + 1
        }
        charIndex[s[end]] = end
        currentLength := end - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}
