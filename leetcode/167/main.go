package main

import "fmt"

func main () {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	println(result)
}	

// Time: O(n)
// Space: O(n)
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)

	// O(n)
	for i, num := range numbers {
		m[num] = i
		// O(1)
		if j, ok := m[target-num]; ok && j != i {
			fmt.Println("j,i", j, i)
			return []int{j+1, i+1}
		}
	}
	return nil
}
