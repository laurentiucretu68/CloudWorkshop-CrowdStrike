package ex1

import (
	"strings"
)

func WordCount(s string) map[string]int {
	freq := map[string]int{}
	words := strings.Split(s, " ")
	for _, word := range words {
		freq[word]++
	}
	return freq
}
