package main

import (
	"math"
)

func scoreOfString(s string) int {
	sum := 0
    for index, char := range s {
		if index == len(s)-1 {
			return sum
		}
		// Rune has the ascii value of the character
		sum += int(math.Abs(float64((char - rune(s[index+1])))))
	}
	return sum
}
