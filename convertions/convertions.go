package convertions

import "errors"

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ConvertRomanToNumber(s string) (int, error) {
	if s == "" {
		return 0, errors.New("Empty value is not valid!")
	}

	total, lastVal := 0, 0

	for i := 0; i < len(s); i++ {

		char := s[len(s)-(i+1) : len(s)-i]
		currentVal := roman[char]

		if lastVal > currentVal {
			total += currentVal
		} else {
			total = total - currentVal
		}

		lastVal = currentVal

	}

	return total * -1, nil
}
