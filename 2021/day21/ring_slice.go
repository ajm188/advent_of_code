package main

import "fmt"

type RingSlice struct {
	b []int
}

func (rs *RingSlice) Get(idx int) int {
	return rs.b[idx%len(rs.b)]
}

func (rs *RingSlice) GetSlice(lo, hi int) []int {
	if hi < lo {
		panic(fmt.Sprintf("cannot GetSlice with hi (%d) < lo (%d)", hi, lo))
	}

	s := make([]int, 0, hi-lo)
	for i := lo; i < hi; i++ {
		s = append(s, rs.Get(i))
	}

	return s
}
