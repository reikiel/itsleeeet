func canConstruct(ransomNote string, magazine string) bool {
    count := [26]int{}
    for _, r := range magazine {
        count[r-'a'] += 1
    }

    for _, r := range ransomNote {
        if count[r-'a'] <= 0 { return false }
        count[r-'a'] -= 1
    }

    return true
}