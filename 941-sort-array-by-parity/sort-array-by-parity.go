func sortArrayByParity(nums []int) []int {
    i := 0

    for j := range nums {
        if nums[j] % 2 == 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    return nums
}