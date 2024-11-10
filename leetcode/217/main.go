package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println(result)
}

// Time complexity, O(n)
// Space complexity, O(n)
func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, num := range nums {
		if m[num] {
			return true
		} else {
			m[num] = true
		}
	}
	return false
}
