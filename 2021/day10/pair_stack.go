package main

import (
	"errors"
	"fmt"
)

var (
	ErrStackEmpty   = errors.New("cannot Pop empty stack")
	ErrStackCorrupt = errors.New("stack is corrupted")
)

type PairStack struct {
	stack   []rune
	closers map[rune]rune
}

func NewPairStack(closers map[rune]rune) *PairStack {
	return &PairStack{
		closers: closers,
	}
}

func (s *PairStack) Len() int {
	return len(s.stack)
}

func (s *PairStack) Push(r rune) error {
	if _, ok := s.closers[r]; ok {
		_, err := s.PopMatch(r)
		return err
	}

	s.stack = append([]rune{r}, s.stack...)
	return nil
}

func (s *PairStack) Peek() (rune, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}

	return s.stack[0], true
}

func (s *PairStack) PopMatch(r rune) (matching rune, err error) {
	if len(s.stack) == 0 {
		return 0, ErrStackEmpty
	}

	if matching, ok := s.closers[r]; ok {
		if matching != s.stack[0] {
			return 0, fmt.Errorf("%w: rune (%s) expects matcher %s, found %s", ErrStackCorrupt, string(r), string(matching), string(s.stack[0]))
		}

		s.stack = s.stack[1:]
		return matching, nil
	}

	return 0, fmt.Errorf("%w: no closer set for %s", ErrStackCorrupt, string(r))
}
