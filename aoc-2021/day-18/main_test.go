package main

import (
	"reflect"
	"testing"
)

var reducePairData = []struct {
	in, out []P
}{
	{
		[]P{P{9,5}, P{8,5}, P{1,4}, P{2,3}, P{3,2}, P{4,1}},
		[]P{P{0,4}, P{9,4}, P{2,3}, P{3,2}, P{4,1}},
	},
	{
		[]P{P{7,1}, P{6,2}, P{5,3}, P{4,4}, P{3,5}, P{2,5}},
		[]P{P{7,1}, P{6,2}, P{5,3}, P{7,4}, P{0,4}},
	},
	{
		[]P{P{6,2}, P{5,3}, P{4,4}, P{3,5}, P{2,5}, P{1,1}},
		[]P{P{6,2}, P{5,3}, P{7,4}, P{0,4}, P{3,1}},
	},
	{
    []P{P{3,2}, P{2,3}, P{1,4}, P{7,5}, P{3,5}, P{6,2}, P{5,3}, P{4,4}, P{3,5}, P{2,5} },
    []P{P{3,2}, P{2,3}, P{8,4}, P{0,4}, P{9,2}, P{5,3}, P{4,4}, P{3,5}, P{2,5} },
  },
	{
    []P{P{3,2}, P{2,3}, P{8,4}, P{0,4}, P{9,2}, P{5,3}, P{4,4}, P{3,5}, P{2,5} },
    []P{P{3,2}, P{2,3}, P{8,4}, P{0,4}, P{9,2}, P{5,3}, P{7,4}, P{0,4}, },
  },
}

func TestExplodePair(t *testing.T) {
	for _, tt := range reducePairData {
		res := Explode(tt.in)
		if !reflect.DeepEqual(res, tt.out) {
			t.Fatalf("Explode(%v)=%v, want %v", tt.in, res, tt.out)
		}
	}
}
