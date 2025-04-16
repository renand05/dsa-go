package arrays

func ContainsDuplicate(nums []int) bool {
	indexedNums := make(map[int]bool)
	for _, num := range nums {
		if _, ok := indexedNums[num]; ok {
			return true
		}
		indexedNums[num] = true
	}
	return false
}
