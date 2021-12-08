package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
	"github.com/ajm188/advent_of_code/pkg/text"
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
	segmentsToNumbers       map[string]int
	uniqueSegmentsToNumbers = map[int]int{}
	allMappings             []map[rune]rune
)

func encode(s string, mapping map[rune]rune) string {
	return strings.Map(func(r rune) rune {
		return mapping[r]
	}, s)
}

func decode(inputs []string, outputs []string) ([]string, []string) {
	var encodings map[string]int
	for _, mapping := range allMappings {
		expectedEncodings := make(map[string]int, len(numberToSegments))
		for k, v := range numberToSegments {
			expectedEncodings[text.SortString(encode(v, mapping))] = k
		}

		found := true
		for _, encdig := range append(inputs, outputs...) {
			if _, ok := expectedEncodings[text.SortString(encdig)]; !ok {
				found = false
				break
			}
		}

		if found {
			encodings = expectedEncodings
			break
		}
	}

	if encodings == nil {
		panic(fmt.Errorf("unable to solve decoding for %v | %v", inputs, outputs))
	}

	decodedInputs := make([]string, len(inputs))
	for i, in := range inputs {
		decodedInputs[i] = numberToSegments[encodings[text.SortString(in)]]
	}
	decodedOutputs := make([]string, len(outputs))
	for i, out := range outputs {
		decodedOutputs[i] = numberToSegments[encodings[text.SortString(out)]]
	}

	return decodedInputs, decodedOutputs
}

func init() {
	segmentsToNumbers = make(map[string]int, len(numberToSegments))
	for k, v := range numberToSegments {
		segmentsToNumbers[v] = k
	}

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

	letters := "abcdefg"
	for _, permutation := range text.Permutations(letters) {
		mapping := make(map[rune]rune, 10)
		for i, r := range letters {
			mapping[r] = rune(permutation[i])
		}

		allMappings = append(allMappings, mapping)
	}
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var part1Count, part2Sum int

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

		for _, encDigit := range strings.Split(outputs, " ") {
			if _, ok := uniqueSegmentsToNumbers[len(encDigit)]; ok {
				part1Count++
			}
		}

		var buf strings.Builder
		_, decodedOutputs := decode(strings.Split(inputs, " "), strings.Split(outputs, " "))
		for _, output := range decodedOutputs {
			fmt.Fprintf(&buf, "%d", segmentsToNumbers[output])
		}

		x, err := strconv.ParseInt(buf.String(), 10, 64)
		cli.ExitOnErrorf(err, "could not parse decoded output %s (line:%d)", buf.String(), i)
		part2Sum += int(x)
	}

	fmt.Println(part1Count)
	fmt.Println(part2Sum)
}
