func removeElement(nums []int, val int) int {
    i := 0
    for _, x := range nums {
        if x != val {
            nums[i] = x
            i++
        }   
    }
    return i
}