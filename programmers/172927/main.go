package main

import (
	"math"
	"sort"
)

func main() {
	picks := []int{1, 0, 1}
	minerals := []string{"iron", "iron", "iron", "iron", "diamond", "diamond", "diamond"}
	result := solution(picks, minerals)
	println(result)
}

func solution(picks []int, minerals []string) int {
	answer := 0
	minCount := min((picks[0]+picks[1]+picks[2])*5, len(minerals))
	minerals = minerals[:minCount]
	groups := make([]map[string]int, int(math.Ceil(float64(minCount)/5.0)))

	dia, iron, stone := picks[0], picks[1], picks[2]

	for i := 0; i < len(minerals); i += 5 {
		groups[i/5] = make(map[string]int)
		for j := 0; j < 5; j++ {
			if i+j < len(minerals) {
				groups[i/5][minerals[i+j]]++
			}
		}
	}
	sort.Slice(groups, func(i, j int) bool {
		return (groups[i]["diamond"] > groups[j]["diamond"]) || (groups[i]["diamond"] == groups[j]["diamond"] && (groups[i]["iron"] > groups[j]["iron"])) || (groups[i]["diamond"] == groups[j]["diamond"] && (groups[i]["iron"] == groups[j]["iron"]) && (groups[i]["stone"] > groups[j]["stone"]))
	})

	for _, group := range groups {
		if dia > 0 {
			for _, value := range group {
				answer += 1 * value
			}
			dia--
		} else if iron > 0 {
			for key, value := range group {
				if key == "diamond" {
					answer += 5 * value
				} else {
					answer += 1 * value
				}
			}
			iron--
		} else if stone > 0 {
			for key, value := range group {
				if key == "diamond" {
					answer += 25 * value
				} else if key == "iron" {
					answer += 5 * value
				} else {
					answer += 1 * value
				}
			}
			stone--
		}
	}

	return answer
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("at least one argument is required")
	}

	minVal := nums[0]
	for _, num := range nums[1:] {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
