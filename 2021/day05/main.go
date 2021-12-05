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
	parseErrMsg := "failed to parse coordinate in %q (line:%d): %s"

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
		cli.ExitOnErrorf(err, parseErrMsg, p1parts[0], i, err)

		y1, err := strconv.ParseInt(p1parts[1], 10, 64)
		cli.ExitOnErrorf(err, parseErrMsg, p1parts[1], i, err)

		p2parts := strings.Split(parts[1], ",")
		if len(p2parts) != 2 {
			cli.ExitOnError(fmt.Errorf(errmsg, line, i))
		}

		x2, err := strconv.ParseInt(p2parts[0], 10, 64)
		cli.ExitOnErrorf(err, parseErrMsg, p2parts[0], i, err)

		y2, err := strconv.ParseInt(p2parts[1], 10, 64)
		cli.ExitOnErrorf(err, parseErrMsg, p2parts[1], i, err)

		vents[i] = &vent{
			P1: &coordinate{X: int(x1), Y: int(y1)},
			P2: &coordinate{X: int(x2), Y: int(y2)},
		}
	}

	hitmap2D := map[string]int{}
	hitmap := map[string]int{}
	for _, vent := range vents {
		xrange := []int{vent.P1.X, vent.P2.X}
		yrange := []int{vent.P1.Y, vent.P2.Y}
		sort.Ints(xrange)
		sort.Ints(yrange)

		if vent.P1.Y == vent.P2.Y {
			for x := xrange[0]; x <= xrange[1]; x++ {
				p := &coordinate{
					X: x,
					Y: vent.P1.Y,
				}

				hitmap2D[p.String()]++
				hitmap[p.String()]++
			}
		} else if vent.P1.X == vent.P2.X {
			for y := yrange[0]; y <= yrange[1]; y++ {
				p := &coordinate{
					X: vent.P1.X,
					Y: y,
				}

				hitmap2D[p.String()]++
				hitmap[p.String()]++
			}
		} else {
			var (
				xstep, ystep func(int) int
				xstop, ystop func(int) bool
			)
			if vent.P1.X < vent.P2.X {
				xstep = func(x int) int { return x + 1 }
				xstop = func(x int) bool { return x <= vent.P2.X }
			} else {
				xstep = func(x int) int { return x - 1 }
				xstop = func(x int) bool { return x >= vent.P2.X }
			}

			if vent.P1.Y < vent.P2.Y {
				ystep = func(y int) int { return y + 1 }
				ystop = func(y int) bool { return y <= vent.P2.Y }
			} else {
				ystep = func(y int) int { return y - 1 }
				ystop = func(y int) bool { return y >= vent.P2.Y }
			}

			for x, y := vent.P1.X, vent.P1.Y; xstop(x) && ystop(y); x, y = xstep(x), ystep(y) {
				p := &coordinate{x, y}
				hitmap[p.String()]++
			}
		}
	}

	var count2D int
	for _, hits := range hitmap2D {
		if hits > 1 {
			count2D++
		}
	}

	var count int
	for _, hits := range hitmap {
		if hits > 1 {
			count++
		}
	}

	fmt.Println(count2D)
	fmt.Println(count)
}
