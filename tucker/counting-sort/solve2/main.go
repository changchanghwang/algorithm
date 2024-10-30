package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 4, 5, 7, 8, 4, 5, 2, 3, 4, 8, 0, 7, 9, 10, 5, 6, 4, 7, 0, 2, 4, 8, 7, 1, 2, 1, 2, 3, 4, 5}

	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	fmt.Println(count)

	sorted := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	fmt.Println(sorted)
}
