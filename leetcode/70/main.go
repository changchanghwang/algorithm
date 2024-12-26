package main

func main() {
	n := 2
	result := climbStairs(n)
	println(result)
}

// Time complexity, O(n)
// Space complexity, O(1)
// 피보나치 수열
func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return b
}
