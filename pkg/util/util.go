package util

import (
	"regexp"
	"strconv"
)

func NumberFromString(str string) (float64, error) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	f := re.FindString(str)
	s, err := strconv.ParseFloat(f, 64)

	return s, err
}

func CreateIndex(data [][]string, idx int) map[string][]string {
	m := make(map[string][]string)
	for _, row := range data {
		if idx < 0 || idx >= len(row) {
			continue
		}

		m[row[idx]] = row
	}

	return m
}
