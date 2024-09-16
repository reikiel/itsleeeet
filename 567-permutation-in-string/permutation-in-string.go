func checkInclusion(s1 string, s2 string) bool {
    window := len(s1)
    
    s1r := []rune(s1)
    sort.Slice(s1r, func(i, j int) bool {
        return s1r[i] < s1r[j]
    })

    for i := 0; i < len(s2)-window+1; i++ {
        curr := s2[i:i+window]
        currr := []rune(curr)
        sort.Slice(currr, func(i, j int) bool {
            return currr[i] < currr[j]
        })
        if string(s1r) == string(currr) { return true }

    }

    return false

}

