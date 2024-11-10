package main

import "sort"

func main() {
	s := "anagram"
	t := "nagaram"
	result := isAnagram(s, t)
	println(result)
	result = isAnagram2(s, t)
	println(result)
}

// Time complexity, O(nlogn)
// Space complexity, O(n)
func isAnagram(s string, t string) bool {
	runes1 := []rune(s)
	runes2 := []rune(t)
	sort.Slice(runes1, func(i, j int) bool { // Quick sort, O(nlogn)
		return runes1[i] < runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] < runes2[j]
	})

	return string(runes1) == string(runes2)
}

// Time complexity, O(n)
// Space complexity, O(1)
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)

	for index, _ := range count {
		count[index] = 0
	}

	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
		count[int(t[i])-int('a')]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}
