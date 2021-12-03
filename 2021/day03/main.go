package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

func pivot(lines []string) []string {
	bufs := make([]strings.Builder, len(lines))
	for _, line := range lines {
		for j, c := range line {
			bufs[j].WriteRune(c)
		}
	}

	result := make([]string, 0, len(bufs))
	for _, buf := range bufs {
		if s := buf.String(); s != "" {
			result = append(result, s)
		}
	}

	return result
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	columns := pivot(lines)
	var (
		mcbs string
		lcbs string
	)

	for _, column := range columns {
		if strings.Count(column, "0") >= len(column)/2 {
			mcbs += "0"
			lcbs += "1"
		} else {
			mcbs += "1"
			lcbs += "0"
		}
	}

	gamma, err := strconv.ParseInt(mcbs, 2, 64)
	cli.ExitOnError(err)
	epsilon, err := strconv.ParseInt(lcbs, 2, 64)
	cli.ExitOnError(err)

	fmt.Printf("gamma=%d\tepsilon=%d\tproduct=%d\n", gamma, epsilon, gamma*epsilon)
}
