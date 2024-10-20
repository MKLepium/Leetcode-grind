package main

func commonChars(words []string) []string {
    seen := make([]int, 26)
	// At max 100 appearances of a character
	for i := range seen {
		seen[i] = 100
	}

	// Counting the characters in each word
	for _, word := range words {
		count := make([]int, 26)
		for _, char := range word {
			count[char-'a']++
		}
		for i, c := range count {
			seen[i] = min(seen[i], c)
		}
	}
	// Creating the result
	res := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < seen[i]; j++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}