func heightChecker(heights []int) int {
    counter := make([]int, 101)

    for _, h := range heights {
        counter[h]++
    }

    res, ctrIdx := 0,0
    
    for _, h := range heights {
        // keep incrementing if 0
        for counter[ctrIdx] == 0 {
            ctrIdx++
        }

        // means h is not supposed to be here
        if h != ctrIdx {
            res++
        }

        counter[ctrIdx]--
    }
    return res
}