package main

import "fmt"

func main () {
	result := DNAStrand("ATTGC")

	fmt.Println(result)
}

var convertMap = map[string]string {
	"A": "T",
	"T": "A",
	"C": "G",
	"G": "C",
}

func DNAStrand(dna string) string {
	var result string

	for _,s := range dna {
		if val, ok := convertMap[string(s)]; ok {
			result += val
		}else {
			result += string(s)
		}

	}

	return result
}