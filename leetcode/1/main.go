package main

import (
	"fmt"
)

func main() {
	result1 := twoSum1([]int{2, 7, 11, 15}, 9)
	fmt.Println(result1)
	result2 := twoSum2([]int{3, 3}, 6)
	fmt.Println(result2)
}

// Time O(n^2)
// Space O(1)
func twoSum1(nums []int, target int) []int {
	// Time: O(n)
	for i := 0; i < len(nums)-1; i++ {
		// Time: O(n)
		for j := i + 1; j < len(nums); j++ {
			// Time: O(1)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Time: O(n)
// Space: O(n)
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	// O(n)
	for i, num := range nums {
		// O(1)
		if j, ok := m[target-num]; ok && j != i {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
