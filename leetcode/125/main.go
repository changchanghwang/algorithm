package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main () {
	s := "0P"
	result := isPalindrome(s)
	fmt.Println(result)
}


//O(n) time complexity
//O(n) space complexity
func isPalindrome(s string) bool {
	regex, _ := regexp.Compile("[a-zA-Z0-9]")
	reverseAndFilteredString := ""
	filteredString := ""

	for _, char := range s {
		if regex.MatchString(string(char)) {
			c := strings.ToLower(string(char))
			reverseAndFilteredString = c  + reverseAndFilteredString
			filteredString +=  c
		}
	}

	return reverseAndFilteredString == filteredString
}

