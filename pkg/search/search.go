package search

import "sort"

func Ints(a []int, x int) (int, bool) {
	idx := sort.SearchInts(a, x)

	if idx == len(a) || a[idx] != x {
		return len(a), false
	}

	return idx, true
}
