package validparentheses

func IsValidParentheses(s string) bool {
	// first in, last out -> stack
	// hold brackets in a stack, check if closing matches the value on the top of the stack
	if len(s) % 2 != 0 {
		return false
	}
	sets := map[string]string{"(": ")", "[": "]", `{`: `}`}
	var stack []string
	for _, ch := range s {
		if _, ok := sets[string(ch)]; !ok && len(stack) == 0 {
			// starts with a closing brace
			return false
		} else if _, ok := sets[string(ch)]; ok {
			// it's an opening brace -> add to stack
			stack = append(stack, string(ch))
		} else {
			// it's a closing brace -> check last value of the stack
			if sets[stack[len(stack)-1]] != string(ch) {
				// if the closing tag doesn't match the last opening tag
				return false
			} else {
				// does match, so pop off the latest opening tag from stack
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
