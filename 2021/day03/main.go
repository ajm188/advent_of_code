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

func partitionBinary(s string, include map[int]struct{}) (zeroes []int, ones []int) {
	keep := func(i int) bool {
		if include == nil {
			return true
		}

		_, ok := include[i]
		return ok
	}

	for i, c := range s {
		if keep(i) {
			switch c {
			case '0':
				zeroes = append(zeroes, i)
			case '1':
				ones = append(ones, i)
			default:
				cli.ExitOnError(fmt.Errorf("invalid binary digit %c (col:%d)", c, i))
			}
		}
	}

	return zeroes, ones
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
		if zeroes, ones := partitionBinary(column, nil); len(zeroes) > len(ones) {
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

	var (
		o2bits  string
		o2rows  = map[int]struct{}{}
		co2bits string
		co2rows = map[int]struct{}{}
	)

	for i := range lines {
		o2rows[i] = struct{}{}
		co2rows[i] = struct{}{}
	}

	for _, column := range columns {
		if len(o2rows) > 1 {
			if zeroes, ones := partitionBinary(column, o2rows); len(ones) >= len(zeroes) {
				for _, skip := range zeroes {
					delete(o2rows, skip)
				}
			} else {
				for _, skip := range ones {
					delete(o2rows, skip)
				}
			}
		}

		if len(co2rows) > 1 {
			if zeroes, ones := partitionBinary(column, co2rows); len(zeroes) <= len(ones) {
				co2bits += "0"
				for _, skip := range ones {
					delete(co2rows, skip)
				}
			} else {
				co2bits += "1"
				for _, skip := range zeroes {
					delete(co2rows, skip)
				}
			}
		}
	}

	for row := range o2rows {
		o2bits = lines[row]
		break
	}

	for row := range co2rows {
		co2bits = lines[row]
		break
	}

	o2, err := strconv.ParseInt(o2bits, 2, 64)
	cli.ExitOnError(err)

	co2, err := strconv.ParseInt(co2bits, 2, 64)
	cli.ExitOnError(err)

	fmt.Printf("o2=%d\tco2=%d\tproduct=%d\n", o2, co2, o2*co2)
}
