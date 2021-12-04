package main

type Board struct {
	spaces  [][]*space
	numbers map[int64]*coordinate
}

func NewBoard(grid [][]int64) *Board {
	b := &Board{
		spaces:  make([][]*space, len(grid)),
		numbers: map[int64]*coordinate{},
	}

	for r, row := range grid {
		spacesRow := make([]*space, len(row))
		for c, val := range row {
			spacesRow[c] = &space{val: val}
			b.numbers[val] = &coordinate{row: r, col: c}
		}

		b.spaces[r] = spacesRow
	}

	return b
}

func (b *Board) Mark(x int64) (ok bool, bingo bool) {
	coord, ok := b.numbers[x]
	if !ok {
		return false, false
	}

	b.spaces[coord.row][coord.col].marked = true

	horizontal := make([]*coordinate, len(b.spaces))
	vertical := make([]*coordinate, len(b.spaces))
	for i := 0; i < len(b.spaces); i++ {
		horizontal[i] = &coordinate{
			row: i,
			col: coord.col,
		}
		vertical[i] = &coordinate{
			row: coord.row,
			col: i,
		}
	}

	bingo = true
	for _, coord := range horizontal {
		if !b.spaces[coord.row][coord.col].marked {
			bingo = false
			break
		}
	}

	if bingo {
		return true, true
	}

	bingo = true
	for _, coord := range vertical {
		if !b.spaces[coord.row][coord.col].marked {
			bingo = false
			break
		}
	}

	return true, bingo
}

func (b *Board) SumUnmarked() (sum int64) {
	for _, row := range b.spaces {
		for _, space := range row {
			if !space.marked {
				sum += space.val
			}
		}
	}

	return sum
}

type space struct {
	marked bool
	val    int64
}

type coordinate struct{ row, col int }
