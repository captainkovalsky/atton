package common

import (
	"errors"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

//Try to parse number from string
func TryParseNumber(s string) (int64, error) {
	numberMatches := re.FindAllString(s, -1)
	for _, element := range numberMatches {
		if s, err := strconv.ParseInt(element, 10, 32); err == nil {
			return s, nil
		} else {
			return 0, errors.New("failed to parse number")
		}
	}

	return 0, errors.New("failed to find any numbers")
}
