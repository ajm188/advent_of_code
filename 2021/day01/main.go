package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

// offsetSumsWhere returns the count of offset sums in xs meeting the condition
// function.
//
// offset sums for a slice are a partition of xs into two slices.
// LHS: [sum(xs[0:n]), sum(xs[1:n+1]), ..., sum(xs[len(xs)-n-1:len(xs)-1])]
// RHS: [sum(xs[1:n]), sum(xs[2:n+2]), ..., sum(xs[len(xs)-n:])]
func offsetSumsWhere(xs []int, n int, f func(l, r int) bool) (count int) {
	l := make([]int, len(xs)-n)
	r := make([]int, len(xs)-n)

	for i, j := 0, 1; i < len(xs)-n && j < len(xs); i, j = i+1, j+1 {
		for k := 0; k < n; k++ {
			l[i] += xs[i+k]
			r[i] += xs[j+k]
		}
	}

	for i := 0; i < len(l); i++ {
		if f(l[i], r[i]) {
			count++
		}
	}

	return count
}

func main() {
	path := flag.String("path", "input.txt", "")
	offsetCSV := flag.String("offsets", "1,3", "csv of integer offsets to sum")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	depths := make([]int, 0, len(lines)-1)

	for i, line := range lines {
		if line == "" {
			continue
		}

		val, err := strconv.ParseInt(line, 10, 64)
		cli.ExitOnErrorf(err, "input line %d", i)

		depths = append(depths, int(val))
	}

	var offsets []int
	for _, offsetStr := range strings.Split(*offsetCSV, ",") {
		offset, err := strconv.ParseInt(offsetStr, 10, 64)
		cli.ExitOnErrorf(err, "could not parse offset csv (%s): %s", *offsetCSV, err)

		offsets = append(offsets, int(offset))
	}

	for _, n := range offsets {
		fmt.Printf("n=%d\t%d\n", n, offsetSumsWhere(depths, n, func(l, r int) bool {
			return l < r
		}))
	}
}
