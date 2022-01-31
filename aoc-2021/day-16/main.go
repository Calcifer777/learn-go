package main

import (
	"fmt"
	"math/big"
	"strconv"
	"utils"
)

func HexToBin(s string) string {
	result := ""
	mappings := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	for _, c := range s {
		result += mappings[c]
	}
	return result
}

func ParseLiteral(s string) (int64, int) {
	value := ""
	var stop int
	for i := 0; i < len(s)+5; i = i + 5 {
		value += s[i+1 : i+5]
		if s[i] == '0' {
			stop = i + 5
			break
		}
	}
	iValue, _ := strconv.ParseInt(value, 2, 64)
	return int64(iValue), stop
}

type Packet struct {
	bits       *string
	value      *big.Int
	typeId     int
	version    int
	subPackets []*Packet
}

func ParsePacket(s string) (Packet, int) {
	version, _ := strconv.ParseInt(s[:3], 2, 32)
	typeId, _ := strconv.ParseInt(s[3:6], 2, 32)
	if typeId == 4 {
		value, stop := ParseLiteral(s[6:])
		p := Packet{
			bits:       &s,
			value:      big.NewInt(value),
			typeId:     int(typeId),
			version:    int(version),
			subPackets: []*Packet{},
		}
		return p, stop + 6
	} else {
		head := s[6]
		if head == '0' {
			subpacketsBits, _ := strconv.ParseInt(s[7:22], 2, 32)
			substr := s[22 : 22+subpacketsBits]
			stop := 22
			subPackets := make([]*Packet, 0)
			for {
				if stop >= 22+int(subpacketsBits) {
					break
				} else {
					pkt, last := ParsePacket(substr)
					stop += last
					substr = substr[last:]
					subPackets = append(subPackets, &pkt)
				}
			}
			p := Packet{
				bits:       &s,
				value:      nil,
				typeId:     int(typeId),
				version:    int(version),
				subPackets: subPackets,
			}
			return p, stop
		} else if head == '1' {
			subpacketNum, _ := strconv.ParseInt(s[7:18], 2, 32)
			stop := 18
			substr := s[18:]
			subPackets := make([]*Packet, 0)
			for i := 0; i < int(subpacketNum); i++ {
				pkt, last := ParsePacket(substr)
				substr = substr[last:]
				stop += last
				subPackets = append(subPackets, &pkt)
			}
			p := Packet{
				bits:       &s,
				value:      nil,
				typeId:     int(typeId),
				version:    int(version),
				subPackets: subPackets,
			}
			return p, stop
		} else {
			panic("Unreachable")
			return Packet{}, 0
		}
	}
}

func SumVersions(p Packet) int {
	if len(p.subPackets) == 0 {
		return p.version
	} else {
		sum := 0
		for _, subpkt := range p.subPackets {
			sum += SumVersions(*subpkt)
		}
		return p.version + sum
	}
}

func (p Packet) Value() *big.Int {
	switch int(p.typeId) {
	case 0: // sum
		{
			value := big.NewInt(0)
			for _, subpkt := range p.subPackets {
				value.Add(value, subpkt.Value())
			}
			return value
		}
	case 1: // prod
		{
			value := big.NewInt(1)
			for _, subpkt := range p.subPackets {
				value.Mul(value, subpkt.Value())
			}
			return value
		}
	case 2: // min
		{
			var value *big.Int = nil
			for _, subpkt := range p.subPackets {
				v := subpkt.Value()
				if value == nil {
					value = v
				} else {
					s := v.Text(10)
					v.Sub(v, value)
					if v.Sign() < 0 {
						tmp, _ := new(big.Int).SetString(s, 0)
						value = tmp
					}
				}
			}
			return value
		}
	case 3: // max
		{
			var value *big.Int = nil
			for _, subpkt := range p.subPackets {
				v := subpkt.Value()
				if value == nil {
					value = v
				} else {
					s := v.Text(10)
					v.Sub(v, value)
					if v.Sign() > 0 {
						tmp, _ := new(big.Int).SetString(s, 10)
						value = tmp
					}
				}
			}
			return value
		}
	case 4: // literal
		{
			return p.value
		}
	case 5: // subpacket_1 > subpacket_2
		{
			v1 := p.subPackets[0].Value()
			v2 := p.subPackets[1].Value()
			diff := v1.Sub(v1, v2)
			if diff.Sign() > 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}
		}
	case 6: // subpacket_1 < subpacket_2
		{
			v1 := p.subPackets[0].Value()
			v2 := p.subPackets[1].Value()
			diff := v1.Sub(v1, v2)
			if diff.Sign() < 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}
		}
	case 7: // subpacket_1 == subpacket_2
		{
			v1 := p.subPackets[0].Value()
			v2 := p.subPackets[1].Value()
			if v1.Cmp(v2) == 0 {
				return big.NewInt(1)
			} else {
				return big.NewInt(0)
			}
		}
	default:
		panic("Value not handled")
	}
}

func Part1(s string) int {
	binary := HexToBin(s)
	p, _ := ParsePacket(binary)
	return SumVersions(p)
}

func Part2(s string) *big.Int {
	binary := HexToBin(s)
	p, _ := ParsePacket(binary)
	return p.Value()
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	utils.Check(err)
	fmt.Printf("Part 1 -> %d\n", Part1(lines[0]))
	value := Part2(lines[0])
	fmt.Printf("Part 2 -> %s\n", value.Text(10))
}
