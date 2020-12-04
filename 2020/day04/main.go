package main

import (
	"flag"
	"fmt"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

func main() {
	path := flag.String("input", "", "path to input, reads from stdin if empty")

	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	passports, err := parse(data)
	cli.ExitOnError(err)

	set := 0

	for _, passport := range passports {
		if passport.IsSet() {
			set++
		}
	}

	fmt.Println(set)
}
