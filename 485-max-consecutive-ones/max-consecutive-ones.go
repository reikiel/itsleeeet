func findMaxConsecutiveOnes(nums []int) int {
    res := 0
    curr := 0
    for _, v := range nums {
        if v == 1 {
            curr += 1
        }
        if v == 0 {
            if curr > res {
                res = curr
            }
            curr = 0
        }
    }

    if curr > res {
        return curr
    }
    return res
}