package main

import "fmt"

func main () {
	fmt.Println(arrayDiff([]int{1,2}, []int{1})) // [2]
}

/**
* a-b
*/
func arrayDiff(a, b []int) []int {
	elementMap := map[int]bool{}
	
	for _, el := range b {
		elementMap[el] = true
	}

	result := []int{}
	for _, v := range a {
		if !elementMap[v] {
			result = append(result,v)
		}
	}

	return result
}