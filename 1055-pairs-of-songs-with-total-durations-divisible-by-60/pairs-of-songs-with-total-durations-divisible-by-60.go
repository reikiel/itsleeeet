func numPairsDivisibleBy60(time []int) int {
    count := 0
    rem := make([]int, 60)

    for _, val := range time {
        r := val % 60

        if r == 0 {
            count += rem[r]
        } else {
            count += rem[60-r]
        }

        rem[r] += 1
    }


    return count
}