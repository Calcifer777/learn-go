package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	out, e := Part1("testdata/sample.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 8
	assert.NotNil(t, out)
	assert.Equal(t, expected, out, "should match")
}

func TestPart1Full(t *testing.T) {
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 7173
	assert.Equal(t, expected, out, "should match")
}

func TestPart2SampleA(t *testing.T) {
	out, e := Part2("testdata/sample-2a.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 4
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Full(t *testing.T) {
	out, e := Part2("testdata/full.txt")
	assert.Nil(t, e, "Part2 failed!")
	expected := 291
	assert.Equal(t, expected, out, "should match")
}
