package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	// moveZeroes(nums)
	moveZeros2(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	zeros := []int{}
	nonZeros := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros = append(zeros, 0)
		} else {
			nonZeros = append(nonZeros, nums[i])
		}
	}

	copy(nums, nonZeros)
	copy(nums[len(nonZeros):], zeros)
}

func moveZeros2(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
	}
}
