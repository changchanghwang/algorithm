package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	fmt.Println(result)
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(start int, remaining int)
	backtrack = func(start int, remaining int) {
		if remaining == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target)
	return result
}
