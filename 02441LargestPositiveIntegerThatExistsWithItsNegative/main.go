package main

import "fmt"

type negpos struct {
	neg bool
	pos bool
}


func findMaxK(nums []int) int {
	seen := make(map[int]negpos)

	for _, val := range nums {
		absVal := max(val, -val)
		neg := false
		if val < 0 {
			neg = true
		}
		
		_, seenBefore := seen[absVal]
		if !seenBefore {
			seen[absVal] = negpos{neg, !neg}
		} else {
			if neg {
				entry := seen[absVal]
				entry.neg = true
				seen[absVal] = entry
			} else {
				entry := seen[absVal]
				entry.pos = true
				seen[absVal] = entry
			}
		}
	}
	largestInt := 0
	for key, val := range seen {
		if val.neg && val.pos {
			fmt.Println(key)
			if largestInt < key {
				largestInt = key
			}
		}
	}

	if largestInt == 0 {
		return -1
	}

	return largestInt
    
}