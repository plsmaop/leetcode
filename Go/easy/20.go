func isValid(s string) bool {
	var stack []rune
	for _, symbol := range s {
		if symbol == '[' || symbol == '(' || symbol == '{' {
			stack = append(stack, symbol)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if stack[len(stack)-1] == '[' && symbol == ']' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '(' && symbol == ')' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '{' && symbol == '}' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}