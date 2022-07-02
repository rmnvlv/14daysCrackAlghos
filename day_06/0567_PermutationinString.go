package day06

import "reflect"

/* Medium */

//Runtime: 1292 ms, Memory Usage: 7.6 MB
func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}

	var createHashMap func(string) map[int32]int

	createHashMap = func(s string) map[int32]int {
		hashMap := make(map[int32]int)
		for _, char := range s {
			hashMap[char] += 1
		}
		return hashMap
	}

	target := createHashMap(s1)

	for i := 0; i <= len(s2)-len(s1); i++ {
		title := createHashMap(s2[i : i+len(s1)])

		if reflect.DeepEqual(title, target) {
			return true
		}
	}

	return false
}
