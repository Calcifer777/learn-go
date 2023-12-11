package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	out, e := Part1("testdata/sample.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := int64(374)
	assert.NotNil(t, out)
	assert.Equal(t, expected, out, "should match")
}

func TestPart1Full(t *testing.T) {
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := int64(9556896)
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Sample(t *testing.T) {
	out, e := Part2("testdata/sample.txt", 10)
	assert.Nil(t, e, "Part2 failed!")
	expected := int64(1030)
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Full(t *testing.T) {
	out, e := Part2("testdata/full.txt", 1000000)
	assert.Nil(t, e, "Part2 failed!")
	expected := int64(685038186836)
	assert.Equal(t, expected, out, "should match")
}
