package arrays

import "sort"

func TopKFrequent(nums []int, k int) []int {
	type Frequency struct {
		Num   int
		Count int
	}

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	frequencies := make([]Frequency, 0, len(counts))
	for num, count := range counts {
		frequencies = append(frequencies, Frequency{Num: num, Count: count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	result := make([]int, 0, k)
	for i := 0; i < k && i < len(frequencies); i++ {
		result = append(result, frequencies[i].Num)
	}

	return result
}
