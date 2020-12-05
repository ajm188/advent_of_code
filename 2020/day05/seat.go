package main

import (
	"bytes"
	"fmt"
)

type Seat struct {
	row int
	col int
}

func (s *Seat) ID() int {
	return s.row*8 + s.col
}

type SeatsByID []*Seat

func (a SeatsByID) Len() int           { return len(a) }
func (a SeatsByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SeatsByID) Less(i, j int) bool { return a[i].ID() < a[j].ID() }

func parseSeats(data []byte) ([]*Seat, error) {
	lines := bytes.Split(data, []byte("\n"))
	seats := make([]*Seat, 0, len(lines))

	for i, line := range lines {
		line := bytes.TrimSpace(line)

		if bytes.Equal(line, []byte{}) {
			continue
		}

		if len(line) != 10 {
			return nil, fmt.Errorf("boarding pass %d was not 10 characters: %s", i, line)
		}

		row, err := coordinate(line[:7])
		if err != nil {
			return nil, err
		}

		col, err := coordinate(line[7:])
		if err != nil {
			return nil, err
		}

		seat := &Seat{
			row: row,
			col: col,
		}
		seats = append(seats, seat)
	}

	return seats, nil
}

func coordinate(pos []byte) (int, error) {
	coord := 0

	for i, x := range pos {
		val := 0

		switch x {
		case 'F', 'f', 'L', 'l':
			val = 0
		case 'B', 'b', 'R', 'r':
			val = 1
		default:
			return -1, fmt.Errorf("invalid rune %c at position %d", x, i)
		}

		place := len(pos) - i - 1
		coord |= (val << place)
	}

	return coord, nil
}
