package arrays_test

import (
	"github.com/renand05/dsa-go/arrays"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: false,
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "one duplicate",
			input:    []int{1, 2, 3, 4, 1},
			expected: true,
		},
		{
			name:     "multiple duplicates",
			input:    []int{1, 2, 3, 4, 1, 2},
			expected: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := arrays.ContainsDuplicate(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
