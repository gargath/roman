package roman

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var test_data = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	846:  "DCCCXLVI",
	900:  "CM",
	1000: "M",
	1999: "MCMXCIX",
	2008: "MMVIII",
}

var _ = Describe("Roman Numerals Converter", func() {
	It("converts numbers correctly", func() {
		for d, r := range test_data {
			Expect(ToRoman(d)).To(Equal(r))
		}
	})
})
