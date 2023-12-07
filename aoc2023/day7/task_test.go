package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Sample(t *testing.T) {
	out, e := Part1("testdata/sample.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := int64(6440)
	assert.NotNil(t, out)
	assert.Equal(t, expected, out, "should match")
}

func TestPart1Full(t *testing.T) {
	out, e := Part1("testdata/full.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := int64(255048101)
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Sample(t *testing.T) {
	out, e := Part2("testdata/sample.txt")
	assert.Nil(t, e, "Part1 failed!")
	expected := int64(5905)
	assert.Equal(t, expected, out, "should match")
}

func TestPart2Full(t *testing.T) {
	out, e := Part2("testdata/full.txt")
	assert.Nil(t, e, "Part2 failed!")
	expected := int64(253718286)
	assert.Equal(t, expected, out, "should match")
}
