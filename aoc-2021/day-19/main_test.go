package main

import (
	"testing"
)

var rotateXTestData = []struct{ in, want V }{
	{V{1, 2, 3}, V{1, -3, 2}},
}

func TestRotateX(t *testing.T) {
	for _, tt := range rotateXTestData {
		out := RotateX(tt.in)
		if out != tt.want {
			t.Fatalf("RotateX(%v)=%v, want %v", tt.in, out, tt.want)
		}
	}
}

var rotateYTestData = []struct{ in, want V }{
	{V{1, 2, 3}, V{3, 2, -1}},
}

func TestRotateY(t *testing.T) {
	for _, tt := range rotateYTestData {
		out := RotateY(tt.in)
		if out != tt.want {
			t.Fatalf("RotateX(%v)=%v, want %v", tt.in, out, tt.want)
		}
	}
}

var rotateZTestData = []struct{ in, want V }{
	{V{1, 2, 3}, V{-2, 1, 3}},
}

func TestRotateZ(t *testing.T) {
	for _, tt := range rotateZTestData {
		out := RotateZ(tt.in)
		if out != tt.want {
			t.Fatalf("RotateX(%v)=%v, want %v", tt.in, out, tt.want)
		}
	}
}

var rotateTestData = []struct {
	in   V
	nx   int
	ny   int
	nz   int
	want V
}{
	{V{1, 2, 3}, 0, 0, 1, V{-2, 1, 3}},
	{V{1, 2, 3}, 0, 0, 2, V{-1, -2, 3}},
	{V{1, 2, 3}, 0, 0, 3, V{2, -1, 3}},
}

func TestRotate(t *testing.T) {
	for _, tt := range rotateTestData {
		out := Rotate(tt.in, tt.nx, tt.ny, tt.nz)
		if out != tt.want {
			t.Fatalf("RotateX(%v)=%v, want %v", tt.in, out, tt.want)
		}
	}
}
