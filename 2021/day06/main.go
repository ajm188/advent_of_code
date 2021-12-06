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
	days := flag.Int("days", 80, "")
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

	counters := strings.Split(line, ",")
	lanternfish := make([]*Lanternfish, len(counters))
	for i, s := range counters {
		counter, err := strconv.ParseInt(s, 10, 64)
		cli.ExitOnErrorf(err, "could not parse counter on line:%d: %w", i, err)

		lanternfish[i] = NewLanternfish(int(counter))
	}

	for i := 0; i < *days; i++ {
		// log.Printf("After %d days: %v", i, nil) // TODO: string representation

		var babies []*Lanternfish
		for _, lf := range lanternfish {
			if baby := lf.Step(); baby != nil {
				babies = append(babies, baby)
			}
		}

		lanternfish = append(lanternfish, babies...)
	}

	fmt.Println(len(lanternfish))
}
