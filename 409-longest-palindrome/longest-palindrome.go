func longestPalindrome(s string) int {
    m := make(map[rune]int)
    res := 0

    for _, r := range s {
        m[r] += 1
    }

    for k,v := range m {
        if v % 2 == 1 {
            res += v-1
            m[k] -= v-1
        } else {
            res += v
            m[k] -= v
        }

        if m[k] == 0 {
            delete(m, k)
        }
    }

    if len(m) != 0 {
        return res + 1
    }

    return res
}