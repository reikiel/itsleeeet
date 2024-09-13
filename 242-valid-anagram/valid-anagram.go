// type sortRunes []rune

// func (s sortRunes) Less(i,j int) bool {
//     return s[i] < s[j]
// }

// func (s sortRunes) Swap(i,j int) {
//     s[i], s[j] = s[j], s[i]
// }

// func (s sortRunes) Len() int {
//     return len(s)
// }

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {return false}

    m := make(map[rune]int)

    for _, v := range s {
        m[v] += 1
    }

    for _,v := range t {
        if _, ok := m[v]; !ok {
            return false
        }
        m[v] -= 1

        if m[v] == 0 { delete(m,v) }
    }

    if len(m) != 0 {return false}

    return true
}