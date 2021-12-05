package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type coordinate struct{ X, Y int }

func (p *coordinate) String() string { return fmt.Sprintf("(%d, %d)", p.X, p.Y) }

type vent struct{ P1, P2 *coordinate }

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	errmsg := "bad input format: %q (line:%d) does not match format 'x1,y1 -> x2,y2'"
	parseErrMsg := "failed to parse coordinate in %q (line:%d): %w"

	vents := make([]*vent, len(lines))
	for i, line := range lines {
		parts := strings.Split(strings.ReplaceAll(line, " ", ""), "->")
		if len(parts) != 2 {
			cli.ExitOnError(fmt.Errorf(errmsg, line, i))
		}

		p1parts := strings.Split(parts[0], ",")
		if len(p1parts) != 2 {
			cli.ExitOnError(fmt.Errorf(errmsg, line, i))
		}

		x1, err := strconv.ParseInt(p1parts[0], 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf(parseErrMsg, p1parts[0], i, err))
		}

		y1, err := strconv.ParseInt(p1parts[1], 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf(parseErrMsg, p1parts[1], i, err))
		}

		p2parts := strings.Split(parts[1], ",")
		if len(p2parts) != 2 {
			cli.ExitOnError(fmt.Errorf(errmsg, line, i))
		}

		x2, err := strconv.ParseInt(p2parts[0], 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf(parseErrMsg, p2parts[0], i, err))
		}

		y2, err := strconv.ParseInt(p2parts[1], 10, 64)
		if err != nil {
			cli.ExitOnError(fmt.Errorf(parseErrMsg, p2parts[1], i, err))
		}

		vents[i] = &vent{
			P1: &coordinate{X: int(x1), Y: int(y1)},
			P2: &coordinate{X: int(x2), Y: int(y2)},
		}
	}

	hitmap := map[string]int{}
	for _, vent := range vents {
		if vent.P1.Y == vent.P2.Y {
			xrange := []int{vent.P1.X, vent.P2.X}
			sort.Ints(xrange)
			for x := xrange[0]; x <= xrange[1]; x++ {
				p := &coordinate{
					X: x,
					Y: vent.P1.Y,
				}

				hitmap[p.String()]++
			}
		}

		if vent.P1.X == vent.P2.X {
			yrange := []int{vent.P1.Y, vent.P2.Y}
			sort.Ints(yrange)
			for y := yrange[0]; y <= yrange[1]; y++ {
				p := &coordinate{
					X: vent.P1.X,
					Y: y,
				}

				hitmap[p.String()]++
			}
		}
	}

	var count int
	for _, hits := range hitmap {
		if hits > 1 {
			count++
		}
	}

	fmt.Println(count)
}
