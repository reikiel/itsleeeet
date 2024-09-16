func check(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
    // No point in processing if length of s1 is greater than s2.
	if n1 > n2 {
		return false
	}

    // Create a map for s1 once. Need a mapping for each character 
    // and its number of occurrences.
	s1Map := make([]int, 26)
	for i := 0; i < n1; i++ {
		s1Map[s1[i]-'a']++
	}

    // Create a similar map for s2. We include all except one character
    // since it will be added in the loop.
	s2Map := make([]int, 26)
	for i := 0; i < n1-1; i++ {
		s2Map[s2[i]-'a']++
	}

    // Sliding window loop. 
	for i := n1 - 1; i < n2; i++ {
        // Add a char from s2 to the window
		s2Map[s2[i]-'a']++

        // Check if s1 and s2 maps are equal. 
        // If they are equal, we have found a permutation.
		if check(s1Map, s2Map) {
			return true
		}

        // Remove the first character from the window to prepare
        // for the next iteration.
		s2Map[s2[i-n1+1]-'a']--
	}

    // No permutation found.
	return false
}