package main

import (
	"math/big"
	"testing"
)

var hexToBinTests = []struct {
	in  string
	out string
}{
	{"38006F45291200", "00111000000000000110111101000101001010010001001000000000"},
	{"D2FE28", "110100101111111000101000"},
}

func TestHexToBin(t *testing.T) {
	for _, tt := range hexToBinTests {
		bin := HexToBin(tt.in)
		if bin != tt.out {
			t.Fatalf(`HexToBin("%s")=%s, want %s`, tt.in, bin, tt.out)
		}
	}
}

var parseLiteralTests = []struct {
	in   string
	out1 int64
	out2 int
}{
	{"10111111100010100000", 2021, 15},
	{"01010", 10, 5},
}

func TestParseLiteral(t *testing.T) {
	for _, tt := range parseLiteralTests {
		value, stop := ParseLiteral(tt.in)
		if value != tt.out1 || stop != tt.out2 {
			t.Fatalf(
				`ParseLiteral("%s")=(%d, %d), want (%d, %d)`,
				tt.in,
				value,
				stop,
				tt.out1,
				tt.out2,
			)
		}
	}
}

var part1Tests = []struct {
	in  string
	out int
}{
	{"38006F45291200", 9},
	{"EE00D40C823060", 14},
	{"8A004A801A8002F478", 16},
	{"620080001611562C8802118E34", 12},
	{"C0015000016115A2E0802F182340", 23},
	{"A0016C880162017C3686B18A3D4780", 31},
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		value := Part1(tt.in)
		if value != tt.out {
			t.Fatalf(`Part1("%s")=%d, want %d`, tt.in, value, tt.out)
		}
	}
}

var part2Tests = []struct {
	in  string
	out *big.Int
}{
	{"C200B40A82", big.NewInt(3)},
	{"04005AC33890", big.NewInt(54)},
	{"880086C3E88112", big.NewInt(7)},
	{"CE00C43D881120", big.NewInt(9)},
	{"D8005AC2A8F0", big.NewInt(1)},
	{"F600BC2D8F", big.NewInt(0)},
	{"9C005AC2F8F0", big.NewInt(0)},
	{"9C0141080250320F1802104A08", big.NewInt(1)},
	{"3232D42BF9400", big.NewInt(5000000000)},
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		value := Part2(tt.in)
		if value.Cmp(tt.out) != 0 {
			t.Fatalf(`Part2("%s")=%d, want %d`, tt.in, value, tt.out)
		}
	}
}
