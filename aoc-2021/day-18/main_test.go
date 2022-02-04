package main

import (
	"reflect"
	"testing"
	"utils"
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
		in := ParseN(tt.in)
		out := ParseN(tt.out)
		res := Explode(in)
		if !reflect.DeepEqual(res, out) {
			t.Fatalf("Input: %v\nout:  %v\nwant: %v", in, res, out)
		}
	}
}

func TestReducePair(t *testing.T) {
	in := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	out := "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"
	res := Reduce(ParseN(in))
	if !reflect.DeepEqual(res, ParseN(out)) {
		t.Fatalf("Reduce(%v)=\n%v\nwant \n%v", in, res, out)
	}
}

var reduceListTestData = []struct {
	in   []string
	want string
}{
	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
	{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}, "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
}

func TestReduceList(t *testing.T) {
	for _, tt := range reduceListTestData {
		in := utils.Map(tt.in, ParseN)
		want := ParseN(tt.want)
		out := ReduceList(in)
		if !reflect.DeepEqual(out, want) {
			t.Fatalf("\nout: %v\nwant: %v", out, tt.want)
		}
	}
}

var magnitudeTestData = []struct {
	in   string
	want int
}{
	{"[[1,2],[[3,4],5]]", 143},
	{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
	{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
	{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
	{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
	{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
}

func TestMagnitude(t *testing.T) {
	for _, tt := range magnitudeTestData {
		in := ParseN(tt.in)
		out := Magnitude(in)
		if out != tt.want {
			t.Fatalf("Magnitude(%v)=%v, want %v", tt.in, out, tt.want)
		}
	}
}
