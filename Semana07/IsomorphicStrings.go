func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sMap := make(map[byte]byte)
    tMap := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        sc := s[i]
        tc := t[i]

        if mappedChar, ok := sMap[sc]; ok {
            if mappedChar != tc {
                return false
            }
        } else {
            sMap[sc] = tc
        }

        if mappedChar, ok := tMap[tc]; ok {
            if mappedChar != sc {
                return false
            }
        } else {
            tMap[tc] = sc
        }
    }

    return true
}
