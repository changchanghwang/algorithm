package main

func main () {

	s := "()}"
	println(isValid(s))

}

// time complexity: O(n)
// space complexity: O(n)
func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '(' || c =='{' || c == '[' {
			stack = append(stack, c)
		}		

		if c == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
		if c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) ==0
}