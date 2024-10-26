package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 4, 5, 7, 8, 4, 5, 2, 3, 4, 8, 0, 7, 9, 10, 5, 6, 4, 7, 0, 2, 4, 8, 7, 1, 2, 1, 2, 3, 4, 5}

	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	sorted := make([]int, 0, len(arr))

	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}

	fmt.Println(sorted)
}
