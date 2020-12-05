package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	seats, err := parseSeats(data)
	cli.ExitOnError(err)

	sort.Sort(sort.Reverse(SeatsByID(seats)))
	fmt.Println(seats[0].ID())
}
