func search(nums []int, target int) int {
    return rec(nums, target, 0, len(nums)-1)
}

func rec(nums []int, target int, l int, r int) int {
    if l > r { return -1 }

    m := l + (r-l)/2

    if nums[m] == target { return m }
    if nums[m] > target { return rec(nums, target, l, m-1) }
    return rec(nums, target, m+1, r)
}