package main

import "fmt"

func main () {
	height := []int{8,7,2,1}
	result := maxArea(height)
	fmt.Println(result)
}

// time complexity: O(n)
// space complexity: O(1)
func maxArea(height []int)int {
	left := 0
	right := len(height)-1
	result := 0
	

	for left < right {
		area := (right-left) * (min(height[left], height[right]))

		if area > result {
			result = area
		}
		if height[left] < height[right]{
			left ++
		}else{
			right --
		}
	}

	return result
}