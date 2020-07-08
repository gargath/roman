package roman

import (
	"strings"
)

var numerals = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

var keys = []int{1000, 500, 100, 50, 10, 5, 1}

// ToRoman converts the integer argument to a string that is its roman numeral
// representation.
func ToRoman(number int) string {
	roman := ""
	for {
		if number == 0 {
			break
		}

		for _, d := range keys {
			for number >= d {
				roman = roman + numerals[d]
				number = number - d
			}
		}
	}

	roman = strings.ReplaceAll(roman, "DCCCC", "CM")
	roman = strings.ReplaceAll(roman, "CCCC", "CD")
	roman = strings.ReplaceAll(roman, "LXXXX", "XC")
	roman = strings.ReplaceAll(roman, "XXXX", "XL")
	roman = strings.ReplaceAll(roman, "VIIII", "IX")
	roman = strings.ReplaceAll(roman, "IIII", "IV")

	return roman
}
