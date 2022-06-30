package day04

import "strings"

/* Easy */

// Runtime: 3 ms, Memory Usage: 6.3 MB
func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		byteWord := []byte(word)

		for first, last := 0, len(word)-1; first < len(word)/2; first, last = first+1, last-1 {
			byteWord[first], byteWord[last] = byteWord[last], byteWord[first]
		}

		words[i] = string(byteWord)
	}

	answer := strings.Join(words, " ")

	return answer

}
