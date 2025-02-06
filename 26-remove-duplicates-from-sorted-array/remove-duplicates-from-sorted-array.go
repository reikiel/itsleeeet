func removeDuplicates(nums []int) int {
    i,j := 0,0
    x := -101
    for j < len(nums) {
        if nums[j] != x {
            nums[i] = nums[j]
            x = nums[j]
            i++
        } 
        j++
    }
    return i
}