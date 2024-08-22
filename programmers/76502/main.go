package main

func main() {
	s := "}]()[{"
	result := solution(s)
	println(result)
}

func solution(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		rotatedS := rotate(s, i)
		if isValid(rotatedS) {
			count++
		}
	}

	return count
}

func rotate(s string, index int) string {
	return s[index:] + s[:index]
}

func isValid(s string) bool {
	stack := []string{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, string(c))
		}
		if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != "[" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
