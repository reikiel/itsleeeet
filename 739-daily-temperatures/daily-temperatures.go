func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []int{}

    for i,v := range temperatures {
        for len(stack) != 0 {
            popI := stack[len(stack)-1]

            if v > temperatures[popI] {
                res[popI] = i-popI
                stack = stack[:len(stack)-1] 
            } else {
                break
            }
        }

        stack = append(stack, i)
    }

    // for len(stack) != 0 {
    //     _, poppedVal := pop(stack)
    //     res[poppedVal] = 0
    // }
    // dont need this cuz alr init to 0

    return res
}

func pop(stack []int) (s []int, val int) {
    val = stack[len(stack)-1]
    s = stack[:len(stack)-1]
    return s, val
}