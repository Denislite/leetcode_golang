package task_20

func isValid(s string) bool {
	bracketsMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	var stack []string

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, string(s[i]))
		} else {
			if len(stack) != 0 {
				if bracketsMap[stack[len(stack)-1]] == string(s[i]) {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
