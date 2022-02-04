package main

import (
	// "fmt"
	"reflect"
	"testing"
	// "utils"
)

var explodeTestData = []struct {
	in, out string
}{
	{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
	{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
	{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
	{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
	{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	{"[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]", "[[[[0,[3,2]],[3,3]],[4,4]],[5,5]]"},
}

func TestExplodePair(t *testing.T) {
	for _, tt := range explodeTestData {
		in := ParseInput(tt.in)
		out := ParseInput(tt.out)
		res := Explode(in)
		if !reflect.DeepEqual(res, out) {
			t.Fatalf("Input: %v\nout:  %v\nwant: %v", in, res, out)
		}
	}
}

func TestReducePair(t *testing.T) {
	in := []P{P{4, 5}, P{3, 5}, P{4, 4}, P{4, 3}, P{7, 3}, P{8, 5}, P{4, 5}, P{9, 4}, P{1, 2}, P{1, 2}}
	out := []P{P{0, 4}, P{7, 4}, P{4, 3}, P{7, 4}, P{8, 4}, P{6, 4}, P{0, 4}, P{8, 2}, P{1, 2}}
	res := Reduce(in)
	if !reflect.DeepEqual(res, out) {
		t.Fatalf("Reduce(%v)=\n%v\nwant \n%v", in, res, out)
	}
}

// var reduceListTestData = []struct {
// 	in   []string
// 	want string
// }{
// 	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
// 	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", }, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
// }
//
// func TestReduceList(t *testing.T) {
// 	for i, tt := range reduceListTestData {
//     fmt.Printf("test RL %d\n", i)
// 		in := utils.Map(tt.in, ParseInput)
// 		want := ParseInput(tt.want)
// 		out := ReduceList(in)
// 		if !reflect.DeepEqual(out, want) {
//       t.Fatalf("\nout: %v\nwant: %v", out, tt.want)
// 		}
// 	}
// }
