package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

func play(boards []*Board, numbers []int64) int64 {
	for _, num := range numbers {
		for _, board := range boards {
			if _, bingo := board.Mark(num); bingo {
				return board.SumUnmarked() * num
			}
		}
	}

	return 0
}

func main() {
	path := flag.String("path", "input.txt", "")
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	lines := strings.Split(string(data), "\n")

	var (
		numbers []int64
		boards  []*Board
	)

	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}

		if strings.Contains(lines[i], ",") {
			if numbers != nil {
				cli.ExitOnError(fmt.Errorf("received two lines of numbers to draw (second set on line:%d)", i))
			}

			nums := strings.Split(lines[i], ",")
			numbers = make([]int64, len(nums))
			for k, s := range nums {
				num, err := strconv.ParseInt(s, 10, 64)
				cli.ExitOnError(err)

				numbers[k] = num
			}

			i++
			continue
		}

		// If it's not empty or a CSV, it's an NxN board.
		var grid [][]int64
		for j := i; j < len(lines) && lines[j] != ""; j++ {
			// sometimes there are extra spaces to pad single-digit numbers
			trimmed := strings.ReplaceAll(strings.TrimLeft(lines[j], " "), "  ", " ")
			rowNums := strings.Split(trimmed, " ")
			row := make([]int64, len(rowNums))
			for k, s := range rowNums {
				num, err := strconv.ParseInt(s, 10, 64)
				cli.ExitOnError(err)

				row[k] = num
			}

			grid = append(grid, row)
		}

		boards = append(boards, NewBoard(grid))

		i += len(grid)
	}

	fmt.Println(play(boards, numbers))
}
