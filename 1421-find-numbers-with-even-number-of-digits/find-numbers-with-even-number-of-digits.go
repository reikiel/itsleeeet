func findNumbers(nums []int) int {
    res := 0
    for _,v := range nums {
        if len(strconv.Itoa(v)) % 2 == 0 {
            res += 1
        }
    }
    return res
}