package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
	"github.com/ajm188/advent_of_code/pkg/sets"
)

var (
	path = flag.String("path", "input.txt", "")

	debug = flag.Bool("debug", false, "")
)

func max(ns ...int) int {
	var m *int
	for _, n := range ns {
		n := n
		if m == nil || n > *m {
			m = &n
		}
	}

	if m == nil {
		return 0
	}

	return *m
}

func min(ns ...int) int {
	var m *int
	for _, n := range ns {
		n := n
		if m == nil || n < *m {
			m = &n
		}
	}

	if m == nil {
		return 0
	}

	return *m
}

type RebootStep struct {
	On     bool
	XRange [2]int
	YRange [2]int
	ZRange [2]int
}

type Coordinate struct {
	X, Y, Z int
}

func (c *Coordinate) String() string { return fmt.Sprintf("(%d, %d, %d)", c.X, c.Y, c.Z) }

func main() {
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var steps []*RebootStep
	rebootRegexp := regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		m := rebootRegexp.FindStringSubmatch(line)
		if m == nil {
			cli.ExitOnError(fmt.Errorf("line %d does not match %s", i, rebootRegexp))
		}

		step := &RebootStep{}
		switch strings.ToLower(m[1]) {
		case "on":
			step.On = true
		case "off":
			step.On = false
		default:
			cli.ExitOnError(fmt.Errorf("invalid on/off value for line %d: %s", i, m[1]))
		}

		parseint := func(s string, name string) int {
			n, err := strconv.ParseInt(s, 10, 64)
			cli.ExitOnErrorf(err, "cannot parse %s for line %d: %s", name, i, err)

			return int(n)
		}

		xlo := parseint(m[2], "xlo")
		xhi := parseint(m[3], "xhi")
		ylo := parseint(m[4], "ylo")
		yhi := parseint(m[5], "yhi")
		zlo := parseint(m[6], "zlo")
		zhi := parseint(m[7], "zhi")

		step.XRange = [2]int{xlo, xhi}
		step.YRange = [2]int{ylo, yhi}
		step.ZRange = [2]int{zlo, zhi}

		steps = append(steps, step)
	}

	if *debug {
		log.Printf("%+v", steps)
	}

	initSet := sets.NewStrings()
	for i, step := range steps {
		if *debug {
			log.Printf("step %d: %+v", i, step)
		}

		f := initSet.Insert
		if !step.On {
			f = initSet.Remove
		}

		for x := max(step.XRange[0], -50); x <= min(step.XRange[1], 50); x++ {
			for y := max(step.YRange[0], -50); y <= min(step.YRange[1], 50); y++ {
				for z := max(step.ZRange[0], -50); z <= min(step.ZRange[1], 50); z++ {
					c := &Coordinate{x, y, z}
					f(c.String())
				}
			}
		}
	}

	fmt.Println(initSet.Len())
}
