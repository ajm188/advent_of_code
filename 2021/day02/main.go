package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type position struct{ x, y int64 }
type aimedPosition struct{ x, y, aim int64 }

type command func(p position, ap aimedPosition) (position, aimedPosition)

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	commands := make([]command, len(lines)-1)

	for i, line := range lines[:len(lines)-1] {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			cli.ExitOnError(fmt.Errorf("input %q (line:%d) did not match format '<command> <units>'", line, i))
		}

		units, err := strconv.ParseInt(parts[1], 10, 64)
		cli.ExitOnErrorf(err, "failed to parse units for input %q (line:%d): %s", line, i, err)

		switch parts[0] {
		case "forward":
			commands[i] = func(p position, ap aimedPosition) (position, aimedPosition) {
				return position{p.x + units, p.y}, aimedPosition{ap.x + units, ap.y + (ap.aim * units), ap.aim}
			}
		case "down":
			commands[i] = func(p position, ap aimedPosition) (position, aimedPosition) {
				return position{p.x, p.y + units}, aimedPosition{ap.x, ap.y, ap.aim + units}
			}
		case "up":
			commands[i] = func(p position, ap aimedPosition) (position, aimedPosition) {
				return position{p.x, p.y - units}, aimedPosition{ap.x, ap.y, ap.aim - units}
			}
		default:
			cli.ExitOnError(fmt.Errorf("invalid command (%s) for input %q (line:%d); must be (forward|down|up)", parts[0], line, i))
		}
	}

	var (
		p  position
		ap aimedPosition
	)
	for _, command := range commands {
		p, ap = command(p, ap)
	}

	fmt.Println(p.x * p.y)
	fmt.Println(ap.x * ap.y)
}
