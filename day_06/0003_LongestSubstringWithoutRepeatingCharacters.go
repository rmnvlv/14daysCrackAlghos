package day06

/* Medium */

//Runtime: 283 ms, Memory Usage: 8.2 MB
func lengthOfLongestSubstring(s string) int {
	answer := 0
	for i := 0; i < len(s); i++ {
		char := make(map[byte]bool)
		char[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := char[s[j]]; ok {
				break
			}
			char[s[j]] = true
		}
		if len(char) > answer {
			answer = len(char)
		}
	}
	return answer
}
