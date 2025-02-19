func moveZeroes(nums []int)  {
    i := 0
    for _, val := range nums {
        if val != 0 {
            nums[i] = val
            i++
        }
    }

    // fill in remaining zeroes
    for ; i < len(nums); i++ {
        nums[i] = 0
    }
}