package main

import "fmt"

func main () {
	result := FindOdd([]int{20 ,1, -1, 2 ,-2 ,3, 3, 5, 5, 1, 2 ,4 ,20, 4 ,-1, -2, 5})
	result2 := FindOdd2([]int{20 ,1, -1, 2 ,-2 ,3, 3, 5, 5, 1, 2 ,4 ,20, 4 ,-1, -2, 5})

	fmt.Println(result)
	fmt.Println(result2)
}

func FindOdd(seq []int) int{
	countMap := map[int]int{}

	for _,num := range seq {
		countMap[num]++
	}

	for k,v := range countMap {
		if v%2 != 0 {
			return k
		}
	}

	return 0
}	
func FindOdd2(seq []int) int {
    res := 0
    for _, x := range seq {
        res ^= x // XOR bit
    }
    return res
}