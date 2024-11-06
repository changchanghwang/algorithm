package main

import (
	"fmt"
	"strings"
)

func main () {
	result := ToJadenCase("How can mirrors be real if our eyes aren't real")

	fmt.Println(result)
}

func ToJadenCase(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if i==0 {
			runes[i] = []rune(strings.ToUpper(string(runes[i])))[0]
		}else if runes[i] == ' ' && i+1 < len(runes) {
			runes[i+1] = []rune(strings.ToUpper(string(runes[i+1])))[0]
		}
	}
	return string(runes)
} 