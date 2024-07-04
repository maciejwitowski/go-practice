package src

import (
	"fmt"
	"strings"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

func GroupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)

	lookup := make(map[string]int)

	for _, str := range strs {
		key := calculateAnagramKey(str)

		index, exists := lookup[key]
		if exists {
			result[index] = append(result[index], str)
		} else {
			lookup[key] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}

func calculateAnagramKey(str string) string {
	count := [26]int{}
	for _, char := range str {
		count[char-'a']++
	}

	var sb strings.Builder
	for i, c := range count {
		if c > 0 {
			sb.WriteString(fmt.Sprintf("%c%d", 'a'+i, c))
		}
	}
	return sb.String()
}
