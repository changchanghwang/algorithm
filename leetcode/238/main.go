package main

import "fmt"

func main () {
	nums := []int{1, 2, 3, 4}
	// result := productExceptSelf(nums)
	result := productExceptSelf(nums)
	fmt.Println(result)
}

// // time complexity: O(n^2)
// // space complexity: O(n)
// func productExceptSelf(nums []int) []int {
// 	result := make([]int, len(nums))

// 	for i := 0; i <len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if result[i] == 0  && j==0{
// 				result[i] = 1
// 			}

// 			if i != j {
// 				result[i] *= nums[j]
// 			}
// 		}
// 	}

// 	return result
// }


// time complexity: O(n)
// space complexity: O(1)
func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    for i := range res {
        res[i] = 1
    }

    prefix := 1
    for i := 0; i < len(nums); i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}