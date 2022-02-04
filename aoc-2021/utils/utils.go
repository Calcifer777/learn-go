package utils


import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(path string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func Min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}

func Max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func Abs(x int) int {
  if x > 0 { 
    return x
  } else {
    return -x
  }
}

func Map[T any, V any](arrT []T, f func(t T) V) []V {
	arrV := make([]V, len(arrT))
	for idx, t := range arrT {
		arrV[idx] = f(t)
	}
	return arrV
}
