package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

type BatteryBank []rune

func (bb BatteryBank) MaxJoltage(n int) int {
	var f func(BatteryBank, int) []rune
	f = func(bb BatteryBank, digits int) []rune {
		if digits == 0 {
			return []rune{}
		}

		var (
			d   rune
			idx int
		)
		for i, b := range bb[0 : len(bb)-digits+1] {
			if b > d {
				d = b
				idx = i
			}
		}

		return append([]rune{d}, f(bb[idx+1:], digits-1)...)
	}

	digits := f(bb, n)
	i, _ := strconv.Atoi(fmt.Sprintf(strings.Repeat("%c", len(digits)), anySlice(digits)...))
	return i
}

func anySlice[T any](slice []T) []any {
	anys := make([]any, len(slice))
	for i, v := range slice {
		anys[i] = v
	}
	return anys
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var banks []BatteryBank
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		banks = append(banks, []rune(line))
	}

	var (
		totalJoltage   int
		totalJoltage12 int
	)
	for _, bank := range banks {
		totalJoltage += bank.MaxJoltage(2)
		totalJoltage12 += bank.MaxJoltage(12)
	}
	fmt.Println(totalJoltage)
	fmt.Println(totalJoltage12)
}
