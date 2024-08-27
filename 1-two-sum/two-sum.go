func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))

    for i,v := range nums {
        if val, ok := m[target-v]; ok {
            return []int{i, val}
        }
        m[v] = i
    }
    return nil
}