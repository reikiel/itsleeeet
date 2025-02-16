func decodeString(s string) string {
    stack := []rune{}

    for _,r := range s{
        if r != ']' {
            stack = append(stack, r)
            continue
        }

        // while not '[', pop
        substr := ""
        for stack[len(stack)-1] != '[' {
            substr = string(stack[len(stack)-1]) + substr
            stack = stack[:len(stack)-1]
        }

        // here you have already got the substr
        // pop the '[' as it is useless now
        stack = stack[:len(stack)-1]

        // now get the number of times to iterate
        // its a rune, and you wont know how mnay digits (1, 2 or 3), so parse it as a string THEN convert
        numStr := ""
        for len(stack)>0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
            numStr = string(stack[len(stack)-1]) + numStr
            stack = stack[:len(stack)-1]
        }
        num, _ := strconv.Atoi(numStr)

        // repeat the string, push it back to the stack
        // need to convert to a 
        substr = strings.Repeat(substr, num)
        stack = append(stack, []rune(substr)...)
    }

    return string(stack)
}