package arrays_test

import (
	"github.com/renand05/dsa-go/arrays"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testcases := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "anagrams",
			s1:   "anagram",
			s2:   "nagaram",
			want: true,
		},
		{
			name: "not anagrams",
			s1:   "rat",
			s2:   "car",
			want: false,
		},
		{
			name: "empty strings",
			s1:   "",
			s2:   "",
			want: true,
		},
		{
			name: "one empty string",
			s1:   "a",
			s2:   "",
			want: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := arrays.IsAnagram(tc.s1, tc.s2)
			if got != tc.want {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v",
					tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
