func thirdMax(nums []int) int {
    mmax, secondMax, thirdMax := math.MinInt64, math.MinInt64, math.MinInt64
    for _,v := range nums {
        if v > mmax {
            thirdMax = secondMax
            secondMax = mmax
            mmax = v
        } else if v < mmax && v > secondMax {
            thirdMax = secondMax
            secondMax = v
        } else if v < secondMax && v > thirdMax {
            thirdMax = v
        }
    }

    if thirdMax == math.MinInt64 {
        return mmax
    } 
    return thirdMax
}