package main

import "fmt"

func main() {
	n := 11
	result := hammingWeight(n)
	println(result)
	result = hammingWeight2(n)
	println(result)
}

func hammingWeight(n int) int {
	count := 0
	binary := fmt.Sprintf("%b", n)
	for _, bit := range binary {
		if bit != '1' {
			continue
		}
		count++
	}

	return count
}

func hammingWeight2(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
