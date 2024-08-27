func isValid(s string) bool {
    // if length is 0 - false - taken care by constraint
    // if length is odd, it will be invalid
    if len(s) % 2 != 0 {return false}

    pairs := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }

    stack := []rune{}

    for _, c := range s {
        if _, ok := pairs[c]; ok { // it is a left bracket
            stack = append(stack, c)
        } else if (len(stack) == 0 || c != pairs[stack[len(stack)-1]] ) { // right bracket, invalid
            return false
        } else { // valid and pop
            stack = stack[:len(stack)-1]
        }
    }

    // if and only if we have empty stack means is valid
    return len(stack) == 0
}