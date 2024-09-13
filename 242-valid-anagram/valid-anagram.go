type sortRunes []rune

func (s sortRunes) Less(i,j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i,j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func isAnagram(s string, t string) bool {
    sr, st := sortRunes(s), sortRunes(t)
    sort.Sort(sr)
    sort.Sort(st)

    return string(sr) == string(st)
}