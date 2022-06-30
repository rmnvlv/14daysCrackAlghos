package day04

/* Easy */

//Runtime: 46 ms, Memory Usage: 6.9 MB
func reverseString(s []byte) {
	first, last := 0, len(s)-1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}

}

//Runtime: 42 ms, Memory Usage: 7 MB
func reverseString2(s []byte) {
	rs := make([]byte, len(s))
	j := len(s) - 1

	for i := 0; i < len(s); i++ {
		rs[j] = s[i]
		j--
	}
	copy(s, rs)
}
