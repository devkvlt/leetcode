package main

import (
	"fmt"
)

var english = map[int]string{
	1000000000: "Billion",
	1000000:    "Million",
	1000:       "Thousand",
	100:        "Hundred",
	90:         "Ninety",
	80:         "Eighty",
	70:         "Seventy",
	60:         "Sixty",
	50:         "Fifty",
	40:         "Forty",
	30:         "Thirty",
	20:         "Twenty",
	19:         "Nineteen",
	18:         "Eighteen",
	17:         "Seventeen",
	16:         "Sixteen",
	15:         "Fifteen",
	14:         "Fourteen",
	13:         "Thirteen",
	12:         "Twelve",
	11:         "Eleven",
	10:         "Ten",
	9:          "Nine",
	8:          "Eight",
	7:          "Seven",
	6:          "Six",
	5:          "Five",
	4:          "Four",
	3:          "Three",
	2:          "Two",
	1:          "One",
}

var parts = []int{
	1000000000,
	1000000,
	1000,
	100,
	90,
	80,
	70,
	60,
	50,
	40,
	30,
	20,
	19,
	18,
	17,
	16,
	15,
	14,
	13,
	12,
	11,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	words := ""

	for num > 0 {
		for _, n := range parts {
			q := num / n
			num %= n
			if q != 0 {
				if q == 1 && n < 100 {
					words += english[n] + " "
				} else {
					words += numberToWords(q) + " " + english[n] + " "
				}
			}
		}
	}

	return words[:len(words)-1]
}

func main() {
	fmt.Printf("|%s|\n", numberToWords(1))
	fmt.Printf("|%s|\n", numberToWords(1994))
	fmt.Printf("|%s|\n", numberToWords(123))
	fmt.Printf("|%s|\n", numberToWords(1234567))
}
