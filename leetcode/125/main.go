package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main () {
	s := "0P"
	result := isPalindrome(s)
	result2 := isPalindrome2(s)
	fmt.Println(result)
	fmt.Println(result2)
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

func isPalindrome2(s string) bool {
	reverseAndFilteredString := ""
	var filteredString strings.Builder
	lowerCaseString := strings.ToLower(s)

	for _, char := range lowerCaseString {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			reverseAndFilteredString = string(char)  + reverseAndFilteredString
			filteredString.WriteRune(char)
		}
	}

	return reverseAndFilteredString == filteredString.String()
}


