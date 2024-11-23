package main

import (
	"fmt"
	"sort"
)

func main () {
	nums := []int{-2,0,0,2,2}
	result := threeSum(nums)
	fmt.Println(result)
}

// time complexity: O(n^2)
// space complexity: O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	for i:=0; i<len(nums); i++ {
	  num := nums[i]
	  if num > 0 {
	 	break
	  }
	  if i > 0 && num == nums[i-1] {
	 	continue
	  }

	  left := i+1
	  right := len(nums)-1


	  for left < right {
		sum :=num + nums[left] + nums[right] 
		if sum > 0 {
			right--
		}else if sum < 0{
			left++
		} else {
			result = append(result, []int{num, nums[left], nums[right]})
			right--
			left++
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		}
	  }
	}

	return result
}