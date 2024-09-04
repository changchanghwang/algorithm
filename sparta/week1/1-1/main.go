package main

import "fmt"

func main() {
	str := "asdfdsa"
	result := isPalindrome(str)
	fmt.Println(result)
}

func isPalindrome(str string) string {
	result := "True"

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			result = "False"
		}
	}

	return result
}
