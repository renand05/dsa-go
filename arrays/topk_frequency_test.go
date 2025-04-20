package arrays

import "testing"

func TestTopKFrequent(t *testing.T) {
	testcases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "basic test",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "k is larger than unique elements",
			nums:     []int{1, 2},
			k:        3,
			expected: []int{1, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := TopKFrequent(tc.nums, tc.k)
			if len(result) != len(tc.expected) {
				t.Errorf("expected length %d, got %d", len(tc.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("expected %d, got %d", tc.expected[i], result[i])
				}
			}
		})
	}
}
