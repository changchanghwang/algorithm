package main

func main () {
	nums := []int{2,7,9,3,1}
	result := rob(nums)
	println(result)
}

func rob(nums []int) int {
	prev := 0
	curr := 0

	for _,num := range nums {
		temp := curr
		curr = max(prev + num, curr)
		prev = temp
	}

	return curr
}