package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("숫자를 입력하세요: ")
	scanner.Scan()
	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("숫자를 입력하세요.")
		return
	}
	result := isOdd(num)
	fmt.Println(result)
}

func isOdd(num int) string {
	if num%2 == 0 {
		return "짝수입니다."
	}
	return "홀수입니다."
}
