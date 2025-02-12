func isValid(s string) bool {
    stack := make([]rune, 0)

    for _,b := range s {
        switch b {
            case ']':
                if invalid(stack,'[') {
                    return false
                }
                stack = stack[:len(stack)-1]  
            case ')':
                if invalid(stack,'(') {
                    return false
                }
                stack = stack[:len(stack)-1]  
            case '}':
                if invalid(stack,'{') {
                    return false
                }
                stack = stack[:len(stack)-1]  
            default:
                stack = append(stack,b)
        }
    }
    
    return len(stack)==0
}

func invalid(stack []rune, bracket rune) bool {
    return len(stack) == 0 || stack[len(stack)-1] != bracket
}