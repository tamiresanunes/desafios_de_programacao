func findSubstring(s string, words []string) []int {
    res := []int{}
    if len(words) == 0 || len(s) == 0 {
        return res
    }

    wordLen := len(words[0])
    wordCount := len(words)
    totalLen := wordLen * wordCount

    wordMap := make(map[string]int)
    for _, word := range words {
        wordMap[word]++
    }

    for i := 0; i <= len(s)-totalLen; i++ {
        seen := make(map[string]int)
        j := 0
        for ; j < wordCount; j++ {
            wordStart := i + j*wordLen
            word := s[wordStart : wordStart+wordLen]
            if count, ok := wordMap[word]; ok {
                seen[word]++
                if seen[word] > count {
                    break
                }
            } else {
                break
            }
        }
        if j == wordCount {
            res = append(res, i)
        }
    }

    return res
}
