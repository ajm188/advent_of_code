package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type point struct{ x, y int }

func (p *point) String() string { return fmt.Sprintf("(%d, %d)", p.x, p.y) }

var (
	ErrBadInput = errors.New("bad input")

	debug = flag.Bool("debug", false, "")
)

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var (
		points []*point
		folds  []Folder

		pointRegexp = regexp.MustCompile(`^(\d+),(\d+)$`)
		foldRegexp  = regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
	)
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		if m := pointRegexp.FindStringSubmatch(line); m != nil {
			if len(m) != 3 {
				cli.ExitOnError(fmt.Errorf("%w: point %q (line %d) does not match %s", ErrBadInput, line, i, pointRegexp))
			}

			x, err := strconv.ParseInt(m[1], 10, 64)
			cli.ExitOnErrorf(err, "%s: cannot parse x-coordinate on line %d: %s", ErrBadInput, i, err)
			y, err := strconv.ParseInt(m[2], 10, 64)
			cli.ExitOnErrorf(err, "%s: cannot parse y-coordinate on line %d: %s", ErrBadInput, i, err)

			points = append(points, &point{x: int(x), y: int(y)})
		} else if m := foldRegexp.FindStringSubmatch(line); m != nil {
			if len(m) != 3 {
				cli.ExitOnError(fmt.Errorf("%w: fold %q (line %d) does not match %s", ErrBadInput, line, i, foldRegexp))
			}

			idx, err := strconv.ParseInt(m[2], 10, 64)
			cli.ExitOnErrorf(err, "%s: cannot parse %s-coordinate on line %d: %s", ErrBadInput, m[1], i, err)
			switch m[1] {
			case "x":
				folds = append(folds, &VerticalFolder{x: int(idx)})
			case "y":
				folds = append(folds, &HorizontalFolder{y: int(idx)})
			default:
				cli.ExitOnError(fmt.Errorf("%w: folds may only be along x or y (got %s on line %d)", ErrBadInput, m[1], i))
			}
		} else {
			cli.ExitOnError(fmt.Errorf("%s is not a point or fold instruction (line %d)", line, i))
		}
	}

	paper := NewPaper(points)
	if *debug {
		fmt.Println(paper.String())
		fmt.Println()
	}
	paper = folds[0].Fold(paper)
	if *debug {
		fmt.Println(paper.String())
		fmt.Println()
	}
}
