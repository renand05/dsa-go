package arrays_test

import (
	"sort"
	"testing"

	"github.com/renand05/dsa-go/arrays"
)

func TestGroupAnagrams(t *testing.T) {
	testcases := []struct {
		name    string
		strings []string
		assert  func(t *testing.T, result [][]string)
	}{
		{
			name:    "empty strings",
			strings: []string{"", ""},
			assert: func(t *testing.T, result [][]string) {
				if len(result) != 1 {
					t.Errorf("expected 1 group, got %d", len(result))
					return
				}
				if len(result[0]) != 2 {
					t.Errorf("expected 2 empty strings in group, got %d", len(result[0]))
					return
				}
				for _, str := range result[0] {
					if str != "" {
						t.Errorf("expected empty string, got %q", str)
					}
				}
			},
		},
		{
			name:    "one group with one anagram, one groups with two anagrams, one group with three anagrams",
			strings: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			assert: func(t *testing.T, result [][]string) {
				for i := range result {
					sort.Strings(result[i])
				}

				sort.Slice(result, func(i, j int) bool {
					return len(result[i]) < len(result[j])
				})

				expected := [][]string{
					{"bat"},
					{"nat", "tan"},
					{"ate", "eat", "tea"},
				}

				for i := range expected {
					sort.Strings(expected[i])
				}

				sort.Slice(expected, func(i, j int) bool {
					return len(expected[i]) < len(expected[j])
				})

				if len(result) != len(expected) {
					t.Fatalf("expected %d groups, got %d", len(expected), len(result))
				}

				for i, group := range result {
					if len(group) != len(expected[i]) {
						t.Errorf(
							"expected group %d to have size %d, got %d",
							i,
							len(expected[i]),
							len(group),
						)
					}
					for j, str := range group {
						if str != expected[i][j] {
							t.Errorf(
								"expected group %d to contain %q, got %q",
								i,
								expected[i][j],
								str,
							)
						}
					}
				}
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := arrays.GroupAnagrams(tc.strings)
			tc.assert(t, result)
		})
	}
}
