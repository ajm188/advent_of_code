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

	mySeatID := -1

	for i := 0; i < len(seats)-1; i++ {
		// higher IDs come first, which is why we subtract in this direction
		if seats[i].ID()-seats[i+1].ID() == 2 {
			mySeatID = seats[i].ID() - 1
			break
		}
	}

	fmt.Println(mySeatID)
}
