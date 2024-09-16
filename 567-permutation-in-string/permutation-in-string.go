

func checkInclusion(s1 string, s2 string) bool {
	
    // No point in processing if length of s1 is greater than s2.
	if len(s1) > len(s2) { return false }

    // keep track of each letter and their count, for each of the string
    // 'b' - 'a' will give 1, and will populate a[1]
    // 'z' - 'a' will give 25, and will populate a[25]
    a1, a2 := [26]int{}, [26]int{}

    // populate s1 letters and counts
    for _, val := range s1 {
        a1[val-'a'] += 1
    }
	

    // For s2, a2 will be the window 
    // we include all except last character
    // since it will be added in the sliding window loop.
    for i := 0; i < len(s1)-1; i++ {
        a2[s2[i]-'a'] += 1
    }
	

    // Sliding window loop
    // add new last character
    // remove old start character 
    start := 0
	for i := len(s1)-1; i < len(s2) ; i++ {
		a2[s2[i]-'a'] += 1

        // Check if the maps are equal. 
        // If they are equal, we have found a permutation.
		if checkEqual(a1, a2) { return true }

        // Remove the first character from the window to prepare
        // for the next iteration.
		a2[s2[start]-'a'] -= 1
        start += 1
	}

    // No permutation found.
	return false
}

func checkEqual(a, b [26]int) bool {
    for i := 0; i < 26 ; i++ {
        if a[i] != b[i] { return false }
    }

    return true
}