package arrays_test

import (
	"github.com/renand05/dsa-go/arrays"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	testcases := []struct {
		name   string
		nums   []int
		wanted []int
	}{
		{
			name:   "basic test",
			nums:   []int{1, 2, 3, 4},
			wanted: []int{24, 12, 8, 6},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := arrays.ProductExceptSelf(tc.nums)
			if len(got) != len(tc.wanted) {
				t.Errorf("expected length %d, got %d", len(tc.wanted), len(got))
				return
			}
			for i := range got {
				if got[i] != tc.wanted[i] {
					t.Errorf("expected %d, got %d", tc.wanted[i], got[i])
				}
			}
		})
	}
}
