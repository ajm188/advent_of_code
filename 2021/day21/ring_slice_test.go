package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rs     []int
		_range [2]int
		out    []int
	}{
		{
			rs:     []int{1, 2, 3, 4, 5},
			_range: [2]int{2, 3},
			out:    []int{3},
		},
		{
			rs:     []int{1, 3, 5, 7, 9, 11, 13},
			_range: [2]int{1, 4},
			out:    []int{3, 5, 7},
		},
		{
			rs:     []int{1, 2, 3},
			_range: [2]int{2, 7},
			out:    []int{3, 1, 2, 3, 1},
		},
		{
			rs:     []int{1},
			_range: [2]int{0, 1},
			out:    []int{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v[%d:%d]", tt.rs, tt._range[0], tt._range[1]), func(t *testing.T) {
			t.Parallel()

			rs := RingSlice{b: tt.rs}
			out := rs.GetSlice(tt._range[0], tt._range[1])
			if !reflect.DeepEqual(out, tt.out) {
				t.Errorf("GetSlice(%v) got %v want %v", tt._range, out, tt.out)
			}
		})
	}
}
