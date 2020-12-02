package main

import (
	"bytes"
	"flag"
	"fmt"
	"sort"
	"strconv"

	"github.com/ajm188/advent_of_code/pkg/cli"
	"github.com/ajm188/advent_of_code/pkg/search"
)

func toIntSlice(data []byte) ([]int, error) {
	lines := bytes.Split(data, []byte("\n"))
	ints := make([]int, 0, len(lines))

	for i, line := range lines {
		if string(line) == "\n" || string(line) == "" {
			continue
		}

		x, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: idx %d", err, i)
		}

		ints = append(ints, int(x))
	}

	return ints, nil
}

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	entries, err := toIntSlice(data)
	cli.ExitOnError(err)

	sort.Ints(entries)

	for _, entry := range entries {
		target := 2020 - entry
		if _, found := search.Ints(entries, target); found {
			fmt.Println(entry * target)
			break
		}
	}

	for i, entry := range entries {
		if i == len(entries)-1 {
			continue
		}

		next := entries[i+1]
		target := 2020 - entry - next

		if _, found := search.Ints(entries, target); found {
			fmt.Println(entry * next * target)
		}
	}
}
