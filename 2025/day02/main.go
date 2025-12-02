package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
	"github.com/ajm188/advent_of_code/pkg/math"
)

type Range struct {
	Start int
	End   int
}

var validIDCache = map[int]bool{}

func (r Range) InvalidIDs() (ids []int) {
	for i := r.Start; i <= r.End; i++ {
		isValid, ok := validIDCache[i]
		if !ok {
			isValid = IsValid(i)
			validIDCache[i] = isValid
		}

		if !isValid {
			ids = append(ids, i)
		}
	}

	return ids
}

func IsValid(id int) bool {
	s := fmt.Sprintf("%d", id)

	if s == strings.Repeat(s[:len(s)/2], 2) {
		return false
	}

	return true
}

func parse(data string) []Range {
	rawRanges := strings.Split(strings.Trim(data, "\n "), ",")
	ranges := make([]Range, len(rawRanges))
	for i, rawRange := range rawRanges {
		parts := strings.Split(rawRange, "-")

		start, err := strconv.Atoi(parts[0])
		cli.ExitOnError(err)

		end, err := strconv.Atoi(parts[1])
		cli.ExitOnError(err)

		ranges[i] = Range{Start: start, End: end}
	}

	return ranges
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	ranges := parse(string(data))

	var allInvalidIDs []int
	for _, r := range ranges {
		ids := r.InvalidIDs()
		allInvalidIDs = append(allInvalidIDs, ids...)
		fmt.Printf("Range %d-%d: %v\n", r.Start, r.End, ids)
	}

	fmt.Printf("Part 1: %d\n", math.Sum(allInvalidIDs))
}
