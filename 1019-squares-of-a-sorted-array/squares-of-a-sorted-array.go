func sortedSquares(nums []int) []int {
    res := make([]int, len(nums))
    var posI int

    // if all positive, then just in order
    if nums[0] >= 0 {
        for i,v := range nums {
            res[i] = v*v
        }
        return res
    }

    // if all is negative, then go backwards
    if nums[len(nums)-1] <= 0 {
        for i,j := len(nums)-1,0; i>=0; {
            res[j] = nums[i]*nums[i]
            i--
            j++
        }
        return res
    }

    // this is the case where theres both positive and negative
    // find the index where it becomes positive
    for i := range nums {
        if nums[i] >= 0 {
            posI = i
            break
        }
    }
    fmt.Println(posI)

    // 2 pointers, for a >= 0 && b < len(nums), a=i-1, b=i
    // if abs(a) > b, put b first 
    a, b, curr := posI-1, posI, 0
    for i:=0; a>=0 && b<len(nums); i++ {
        if abs(nums[a]) > nums[b] {
            res[i] = nums[b] * nums[b]
            b++
        } else {
            res[i] = nums[a] * nums[a]
            a--
        }
        curr = i+1
    }

    // at the end check if a != 0 
    if a >= 0 {
        for i:=a; i>=0; i-- {
            res[curr] = nums[i] * nums[i]
            curr++
        }
    }

    // separate check if b != len(nums)
    if b <= len(nums)-1 {
        for i:=b; i<len(nums); i++ {
            res[curr] = nums[i] * nums[i]
            curr++
        }
    }
    return res
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}