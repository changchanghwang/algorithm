package main

import "sort"

func main() {
	nums := []int{0, 1, 0, 3, 2, 3}
	result := lengthOfLIS(nums)
	println(result)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sub := []int{} // 증가 부분 수열을 추적할 배열

	for _, num := range nums {
		i := sort.Search(len(sub), func(i int) bool { return sub[i] >= num })

		// sub 배열에서 적절한 위치에 num을 삽입하거나 대체
		if i < len(sub) {
			sub[i] = num
		} else {
			sub = append(sub, num)
		}
	}

	return len(sub) // sub 배열의 길이가 최장 증가 부분 수열의 길이
}
