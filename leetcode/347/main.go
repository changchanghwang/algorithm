package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}

// Time: O(nlogn)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}

	result := [][]int{}

	for key, value := range hashMap {
		result = append(result, []int{key, value})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})

	resultNums := []int{}
	for i := 0; i < k; i++ {
		resultNums = append(resultNums, result[i][0])
	}

	return resultNums[:k]
}
