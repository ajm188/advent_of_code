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
	debug := flag.Bool("debug", false, "print each day of the simulation; useful for working out the math (WARNING: slow for -days values starting around 130)")
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

	if *debug {
		for d := 0; d < *days; d++ {
			fmt.Printf("%0.2d days left\t%v\n", *days-d, lanternfish)
			var babies []*Lanternfish
			for _, lf := range lanternfish {
				if baby := lf.Step(); baby != nil {
					babies = append(babies, baby)
				}
			}

			lanternfish = append(lanternfish, babies...)
		}

		fmt.Printf("00 days left\t%v\n", lanternfish)
		return
	}

	// spawnValues maps the days remaining when a fish spawns to the number of
	// fish it will spawn.
	spawnValues := map[int]int{}
	for d := 0; d <= *days+8; d++ {
		// A fish that spawns with less than 8 days left spawns 0 additional
		// fish, because the counter will go:
		//	8 => 7 => 6 => 5 => 4 => 3 => 2 => 1 => (SPAWN) 0
		//
		// That is, on the 8th day, you spawn another fish.
		if d <= 7 {
			spawnValues[d] = 0
			continue
		}

		var value int
		// A fish first spawns another after 8 days (d-8); then for each spawn
		// after it takes 7 days [because 0 is included in the range] (i-=7).
		for i := d - 8; i > 0; i -= 7 {
			// When a new fish is spawned, it does not actually start ticking in
			// the simulation until the day following (i-1).
			value += 1 + spawnValues[i-1]
		}

		spawnValues[d] = value
	}

	var value int
	for _, lf := range lanternfish {
		spawnDate := *days + (8 - lf.counter)
		value += spawnValues[spawnDate] + 1
	}

	fmt.Println(value)
}
