package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	out, e := Part1("testdata/sample.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := 136
	assert.NotNil(t, out)
	assert.Equal(t, expected, out, "should match")
}

func TestPart1Full(t *testing.T) {
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := -1
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Sample(t *testing.T) {
	t.Skip("skip")
	out, e := Part2("testdata/sample.txt")
	assert.Nil(t, e, "Part2 failed!")
	expected := 46
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Full(t *testing.T) {
	t.Skip("skip")
	out, e := Part2("testdata/full.txt")
	assert.Nil(t, e, "Part2 failed!")
	expected := -1
	assert.Equal(t, expected, out, "should match")
}
