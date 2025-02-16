func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); {
        if nums[i] == 0{
            i++
            continue
        }
        if nums[nums[i] - 1] == 0 {
            nums[i] = i + 1
            i++
            continue
        }
        
        nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], 0
    }
    
    j := 0
    for i := 0; i < len(nums); i++ { 
        if nums[i] != 0 {
            nums[j] = nums[i]
            j++
        }
    }
    
    return nums[:j]
}