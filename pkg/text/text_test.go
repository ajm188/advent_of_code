package text

import (
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   string
		outs []string
	}{
		{
			in: "abc",
			outs: []string{
				"abc",
				"acb",
				"bac",
				"bca",
				"cab",
				"cba",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			outs := Permutations(tt.in)

			sort.Strings(outs)
			sort.Strings(tt.outs)

			if len(outs) != len(tt.outs) {
				t.Fatalf("Permutations(%s) len mismatch; want %d permutations, got %d", tt.in, len(tt.outs), len(outs))
			}

			for i, p := range outs {
				if p != tt.outs[i] {
					t.Errorf("Permutations(%s):%d got %v want %v", tt.in, i, p, tt.outs[i])
				}
			}
		})
	}
}

func TestSortString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in  string
		out string
	}{
		{
			"hello",
			"ehllo",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			out := SortString(tt.in)
			if out != tt.out {
				t.Errorf("SortString(%s) got %v want %v", tt.in, out, tt.out)
			}
		})
	}
}
