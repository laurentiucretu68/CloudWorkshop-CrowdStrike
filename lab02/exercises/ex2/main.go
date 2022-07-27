package ex2

import (
	"strings"
)

func VowelsCount(s []string) map[string]int {
	vowelsFreq := map[string]int{}
	vowels := "aeiou"
	for i := range s {
		word := s[i]
		for l := range word {
			if strings.Contains(vowels, string(word[l])) {
				vowelsFreq[word]++
			}
		}
	}
	return vowelsFreq
}
