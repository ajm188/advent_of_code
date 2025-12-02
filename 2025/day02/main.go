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

var (
	validIDCache1 = map[int]bool{}
	validIDCache2 = map[int]bool{}
)

func (r Range) InvalidIDs1() (ids []int) {
	for i := r.Start; i <= r.End; i++ {
		isValid, ok := validIDCache1[i]
		if !ok {
			isValid = IsValid1(i)
			validIDCache1[i] = isValid
		}

		if !isValid {
			ids = append(ids, i)
		}
	}

	return ids
}

func (r Range) InvalidIDs2() (ids []int) {
	for i := r.Start; i <= r.End; i++ {
		isValid, ok := validIDCache2[i]
		if !ok {
			isValid = IsValid2(i)
			validIDCache2[i] = isValid
		}

		if !isValid {
			ids = append(ids, i)
		}
	}

	return ids
}

func IsValid1(id int) bool {
	s := fmt.Sprintf("%d", id)

	if s == strings.Repeat(s[:len(s)/2], 2) {
		return false
	}

	return true
}

func IsValid2(id int) bool {
	s := fmt.Sprintf("%d", id)

	for i := 0; i < len(s)/2; i++ {
		seq := s[0 : i+1]
		n := len(s) / len(seq)
		if strings.Repeat(seq, n) == s {
			return false
		}
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

	var (
		allInvalidIDsPart1 []int
		allInvalidIDsPart2 []int
	)

	for _, r := range ranges {
		ids := r.InvalidIDs1()
		allInvalidIDsPart1 = append(allInvalidIDsPart1, ids...)
		fmt.Printf("Range %d-%d: %v\n", r.Start, r.End, ids)

		ids = r.InvalidIDs2()
		allInvalidIDsPart2 = append(allInvalidIDsPart2, ids...)
		fmt.Printf("Range %d-%d: %v\n", r.Start, r.End, ids)
	}

	fmt.Printf("Part 1: %d\n", math.Sum(allInvalidIDsPart1))
	fmt.Printf("Part 2: %d\n", math.Sum(allInvalidIDsPart2))
}
