func findDisappearedNumbers(nums []int) []int {
    for i:=0;i<len(nums); {
        if nums[i] == 0 {
            i++
            continue
        }
        if nums[nums[i]-1] == 0 || nums[i] == nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = i+1, 0
            i++
            continue
        }

        nums[i], nums[nums[i]-1] = nums[nums[i]-1], 0
    }

    j := 0
    for _,v := range nums {
        if v != 0 {
            nums[j] = v
            j++
        }
    }
    return nums[:j]
}