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
	stack []rune

	pairs        map[rune]rune
	reversePairs map[rune]rune
}

func NewPairStack(pairs []string) *PairStack {
	pairsMap := make(map[rune]rune, len(pairs))
	pairsRev := make(map[rune]rune, len(pairs))

	for _, pair := range pairs {
		if len(pair) != 2 {
			panic("cannot initialize PairStack with 'pair' not length 2, got: " + pair)
		}

		o, c := rune(pair[0]), rune(pair[1])
		pairsMap[o] = c
		pairsRev[c] = o
	}

	return &PairStack{
		pairs:        pairsMap,
		reversePairs: pairsRev,
	}
}

func (s *PairStack) Len() int {
	return len(s.stack)
}

func (s *PairStack) Push(r rune) (popped rune, pushed bool, err error) {
	if _, ok := s.pairs[r]; ok {
		s.stack = append([]rune{r}, s.stack...)
		return 0, true, nil
	}

	if match, ok := s.reversePairs[r]; ok {
		if len(s.stack) == 0 {
			return 0, false, fmt.Errorf("%w: cannot pop %s (resulted from pushing %s)", ErrStackEmpty, string(match), string(r))
		}

		if s.stack[0] != match {
			return 0, false, fmt.Errorf("%w: top of stack (%s) does not match %s (resulted from pushing %s)", ErrStackCorrupt, string(s.stack[0]), string(match), string(r))
		}

		r2 := s.stack[0]
		s.stack = s.stack[1:]
		return r2, false, nil
	}

	return 0, false, fmt.Errorf("%w: unknown character %s", ErrStackCorrupt, string(r))
}

func (s *PairStack) Peek() (rune, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}

	return s.stack[0], true
}
