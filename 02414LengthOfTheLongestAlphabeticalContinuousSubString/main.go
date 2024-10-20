package main

func longestContinuousSubstring(s string) int {
	maxLen := 0
	currentLen := 0
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			currentLen++
		} else {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 1
		}
	}
	if currentLen > maxLen {
		maxLen = currentLen
	}
	return maxLen
}