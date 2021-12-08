package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

var (
	numberToSegments = map[int]string{
		0: "abcefg",
		1: "cf",
		2: "acdeg",
		3: "acdfg",
		4: "bcdf",
		5: "abdfg",
		6: "abdefg",
		7: "acf",
		8: "abcdefg",
		9: "abcdfg",
	}
	uniqueSegmentsToNumbers = map[int]int{}
)

func init() {
	var dupes []int
	for k, v := range numberToSegments {
		key := len(v)
		if _, ok := uniqueSegmentsToNumbers[key]; ok {
			dupes = append(dupes, key)
		}

		uniqueSegmentsToNumbers[key] = k
	}
	for _, dupe := range dupes {
		delete(uniqueSegmentsToNumbers, dupe)
	}
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var part1Count int

	// NOTE: each line has a different encoding, which was very confusing to me
	// at first.
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			cli.ExitOnError(fmt.Errorf("malformed input on line %d: does not match '<digit> <digit> ... <digit> | <digit> ... <digit>", i))
		}

		inputs := strings.TrimSpace(parts[0])
		outputs := strings.TrimSpace(parts[1])

		_ = inputs // not needed for part1
		for _, encDigit := range strings.Split(outputs, " ") {
			if _, ok := uniqueSegmentsToNumbers[len(encDigit)]; ok {
				part1Count++
			}
		}
	}

	fmt.Println(part1Count)
}
