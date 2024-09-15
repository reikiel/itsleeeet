func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    sum := 0
    res := make([]int, 0)

    // count current even sum
    for _, v := range nums {
        if isEven(v) {
            sum += v
        }
    }

    fmt.Println(sum)

    for _, q := range queries {
        if isEven(nums[q[1]]) {
            sum -= nums[q[1]]
        }

        newVal := nums[q[1]] + q[0]
        
        if isEven(newVal) {
            sum += newVal
        }

        nums[q[1]] = newVal

        res = append(res, sum)
    }

    return res
}

func isEven(num int) bool {
    return num % 2 == 0
}