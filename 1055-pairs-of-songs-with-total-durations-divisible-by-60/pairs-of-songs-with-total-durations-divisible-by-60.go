func numPairsDivisibleBy60(time []int) int {
    res:=0
    m := [60]int{}

    for _, v := range time {
        rem := v % 60

        if rem == 0 {
            res += m[0]
        } else {
            res += m[60-rem]
        }

        m[rem] = m[rem] + 1

    } 
    
    return res
}