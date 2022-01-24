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
