package day7

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func Part1(path string) (int64, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	hands, e := parseFile(f)
	if e != nil {
		panic(e)
	}
	sort.Sort(Hands(hands))
	var v int64 = 0
	for idx, h := range hands {
		slog.Info("Part1",
			slog.String("H", h.ToString()),
			slog.String("V", h.Value().String()),
		)
		v += int64((idx + 1) * h.bid)
	}
	slog.Info("Part1",
		slog.Any("result", v),
	)
	return v, nil
}

func Part2(path string) (int, error) {
	f, e := os.Open(path)
	if e != nil {
		slog.Error(fmt.Sprintf("Cound not open file at %s", path))
		return -1, e
	}
	defer f.Close()
	parseFile(f)
	return -1, nil
}

func parseFile(f *os.File) ([]Hand, error) {
	buf := bufio.NewScanner(f)
	hands := make([]Hand, 0)
	for buf.Scan() {
		line := buf.Text()
		hand, e := parseHand(line)
		if e != nil {
			slog.Error("Error parsing line")
			panic(e)
		}
		hands = append(hands, *hand)
		slog.Debug("parsefile",
			slog.String("line", line),
			slog.String("hand", hand.ToString()),
			slog.String("hand value", hand.Value().String()),
		)

	}
	return hands, nil
}

type Hand struct {
	cards     []Card
	bid       int
	cardFreqs map[Card]int
}

func (h Hand) ToString() string {
	return fmt.Sprintf("Hand(cards=%v, bid=%d)", h.cards, h.bid)
}

func NewHand(cards []Card, bid int) Hand {
	return Hand{
		cards:     cards,
		bid:       bid,
		cardFreqs: nil,
	}
}

func (h Hand) getCardsFreqs() map[Card]int {
	if h.cardFreqs == nil {
		freqs := make(map[Card]int)
		for _, c := range h.cards {
			freqs[c] += 1
		}
		h.cardFreqs = freqs
	}
	return h.cardFreqs
}

func maxFreq[T comparable](freqs map[T]int) (T, int) {
	var maxK T
	maxV := -1
	for v, freq := range freqs {
		if freq > maxV {
			maxV = freq
			maxK = v
		}
	}
	return maxK, maxV
}

func (h Hand) Value() HandValue {
	maxFreqCard, f := maxFreq(h.getCardsFreqs())
	slog.Debug("hand.value",
		slog.Int("maxFreq", f),
	)
	switch f {
	case 1:
		return HighCard
	case 2:
		{
			snd := -1
			for c, f := range h.getCardsFreqs() {
				if c == maxFreqCard {
					continue
				} else if f > snd {
					snd = f
				}
			}
			if snd == 2 {
				return TwoPairs
			} else {
				return Pair
			}
		}
	case 3:
		{
			snd := -1
			for c, f := range h.getCardsFreqs() {
				if c == maxFreqCard {
					continue
				} else if f > snd {
					snd = f
				}
			}
			if snd == 1 {
				return ThreeOfAKind
			} else {
				return FullHouse
			}
		}
	case 4:
		return FourOfAKind
	case 5:
		return FiveOfAKind
	default:
		fmt.Errorf("Too high max card frequency!")
	}
	return -1
}

type Hands []Hand

func (hs Hands) Len() int {
	return len(hs)
}

func (hs Hands) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs Hands) Less(i, j int) bool {
	if hs[i].Value() < hs[j].Value() {
		return true
	} else if hs[i].Value() > hs[j].Value() {
		return false
	} else {
		var cI, cJ Card
		for cardIdx := 0; cardIdx < 5; cardIdx++ {
			cI = hs[i].cards[cardIdx]
			cJ = hs[j].cards[cardIdx]
			if cI < cJ {
				return true
			} else if cI > cJ {
				return false
			}
		}
		return false
	}
}

func parseHand(s string) (*Hand, error) {
	chunks := strings.Split(s, " ")
	if len(chunks) != 2 {
		return nil, fmt.Errorf("Error parsing hand: %s", s)
	}
	cards := make([]Card, 0)
	var v Card
	for _, c := range chunks[0] {
		if c == 'A' {
			v = A
		} else if c == 'K' {
			v = K
		} else if c == 'Q' {
			v = Q
		} else if c == 'J' {
			v = J
		} else if c == 'T' {
			v = T
		} else if unicode.IsDigit(c) {
			i, _ := strconv.Atoi(string(c))
			v = Card(i)
		} else {
			return nil, fmt.Errorf("Could not get parse card: %v", c)
		}
		cards = append(cards, v)
	}
	//
	bid, e := strconv.Atoi(chunks[1])
	if e != nil {
		panic(e)
	}
	h := NewHand(cards, bid)
	return &h, nil
}

type Card int

const (
	Two Card = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	J
	Q
	K
	A
)

func (c Card) String() string {
	return [...]string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}[c-2]
}

type HandValue int

const (
	HighCard HandValue = iota
	Pair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (hv HandValue) String() string {
	return [...]string{
		"HighCard",
		"Pair",
		"TwoPairs",
		"ThreeOfAKind",
		"FullHouse",
		"FourOfAKind",
		"FiveOfAKind",
	}[hv]
}
