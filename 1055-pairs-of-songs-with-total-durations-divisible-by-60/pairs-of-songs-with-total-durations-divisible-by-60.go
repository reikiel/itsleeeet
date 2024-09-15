func numPairsDivisibleBy60(time []int) int {
    res:=0
    record:=map[int]int{}
    for _,d:=range time{
        d=d%60

        if d == 0 {
            res += record[d]
        } else {
            res += record[60-d]
        }


        record[d] += 1
    }
    return res
}