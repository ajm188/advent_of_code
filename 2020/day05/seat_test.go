package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinate(t *testing.T) {
	tests := []struct {
		pass     string
		expected int
	}{
		{
			// 1000110
			pass:     "BFFFBBF",
			expected: 70,
		},
		{
			pass:     "FFFBBBF",
			expected: 14,
		},
		{
			pass:     "BBFFBBF",
			expected: 102,
		},
		{
			pass:     "RRR",
			expected: 7,
		},
		{
			pass:     "RLL",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, err := coordinate([]byte(tt.pass))
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
