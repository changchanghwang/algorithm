package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)


	result2 := groupAnagrams2(strs)
	fmt.Println(result2)
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

// Time Complexity O(n * klogk)
// Space Complexity O(n)
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)

		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}


func getKey(s string) string {
    counts := make([]int, 26) // 알파벳 소문자만 있다고 가정
    
    // 각 문자의 빈도수 계산 O(k)
    for _, ch := range s {
        counts[ch-'a']++
    }
    
    // 빈도수 배열을 문자열로 변환 O(1) - 항상 26개
    var key strings.Builder
    for i := 0; i < 26; i++ {
        // '#'를 구분자로 사용하여 같은 빈도수라도 다른 문자를 구분
        key.WriteString(fmt.Sprintf("#%d", counts[i]))
    }
    
    return key.String()
}

func groupAnagrams2(strs []string) [][]string {
    anagrams := make(map[string][]string)
    
    // O(n * k)
    for _, str := range strs {
        key := getKey(str)
        anagrams[key] = append(anagrams[key], str)
    }
    
    // O(n)
    result := make([][]string, 0, len(anagrams))
    for _, group := range anagrams {
        result = append(result, group)
    }
    
    return result
}