package main

import "strconv"

func main () {

	result := []string{"4","13","5","/","+"}
	println(evalRPN(result))
}

func evalRPN(tokens []string) int {
	stack := []int{}
	
	for _, token := range tokens {
		if token == "+" {
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "-" {
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "*" {
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if token == "/" {
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}	
	}
	return stack[0]
}