package main

import (
	"flag"
	"fmt"
	"sort"
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
	var line string
	for i, l := range lines {
		if strings.Contains(l, ",") {
			if line != "" {
				cli.ExitOnError(fmt.Errorf("bad input: only one CSV line allowed; second found on line:%d", i))
			}

			line = l
		}
	}

	var positions []int
	for i, s := range strings.Split(line, ",") {
		pos, err := strconv.ParseInt(s, 10, 64)
		cli.ExitOnErrorf(err, "could not parse position on col:%d: %w", i, err)

		positions = append(positions, int(pos))
	}

	sort.Ints(positions)

	var (
		best *int
		prev *int
	)

	for _, target := range positions {
		if prev != nil && *prev == target {
			continue
		}

		var cost int
		for _, crab := range positions {
			if target == crab {
				continue
			}

			unitCost := target - crab
			if unitCost < 0 {
				unitCost *= -1
			}

			cost += unitCost
		}

		if best == nil || *best >= cost {
			best = &cost
		}
	}

	fmt.Println(*best)
}
