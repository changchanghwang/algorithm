package main

func main() {
	n := 2
	result := climbStairs(n)
	println(result)
}

func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return b
}
