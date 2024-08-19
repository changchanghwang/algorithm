package main

import (
	"fmt"
)

func main() {
	num := 5
	result := countingBits(num)

	fmt.Println(result)
}

func countingBits(num int) []int {
	result := make([]int, num+1)

	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1 // 가장 오른쪽 1을 뺀 비트카운트
		// result[i] = bits.OnesCount(uint(i)) // 내장함수 사용
	}

	return result
}
