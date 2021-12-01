package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	depths := make([]int64, 0, len(lines)-1)

	for i, line := range lines {
		if line == "" {
			continue
		}

		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf("input line %d: %w", i, err))
		}

		depths = append(depths, val)
	}

	var numIncreases int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			numIncreases++
		}
	}

	fmt.Printf("%d\n", numIncreases)
}
