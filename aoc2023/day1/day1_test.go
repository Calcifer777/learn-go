package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	path := "testdata/full.txt"
	out := Part1(path)
	assert.Equal(t, 54916, out, "they should be equal")
}

func TestPart2(t *testing.T) {
	path := "testdata/full.txt"
	out := Part2(path)
	assert.Equal(t, 54728, out, "they should be equal")
}
