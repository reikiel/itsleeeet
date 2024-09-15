func numPairsDivisibleBy60(time []int) int {
    res:=0
    record:=map[int]int{}
    for _,d:=range time{
        d=d%60
        target:=60-d

        if target == 60 {
            target = 0
        }
        
        if num,ok:=record[target];ok{
            res+=num
        }

        record[d] += 1
    }
    return res
}