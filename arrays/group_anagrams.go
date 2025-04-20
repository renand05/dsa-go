package arrays

import (
	"sort"
)

func GroupAnagrams(strs []string) [][]string {
	anagramsMap := make(map[string][]string)

	for _, s := range strs {
		sortedChar := sortChars([]rune(s))
		anagramMapkey := string(sortedChar)
		anagramsMap[anagramMapkey] = append(anagramsMap[anagramMapkey], s)
	}

	result := [][]string{}
	for _, group := range anagramsMap {
		result = append(result, group)
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})

	return result
}

func sortChars(chars []rune) []rune {
	n := len(chars)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if chars[j] > chars[j+1] {
				chars[j], chars[j+1] = chars[j+1], chars[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return chars
}
