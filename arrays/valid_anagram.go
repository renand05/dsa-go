package arrays

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)
	for _, char := range s1 {
		counts[char]++
	}
	for _, char := range s2 {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}
