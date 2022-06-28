package day02

/* Easy */

// Runtime: 30ms, Memory Usage: 6.7mb
func sortedSquares(nums []int) []int {
	len := len(nums)
	first, last := 0, len-1

	answer := make([]int, len)

	for i := len - 1; i >= 0; i-- {
		if abs(nums[first]) > abs(nums[last]) {
			answer[i] = nums[first] * nums[first]
			first++
		} else {
			answer[i] = nums[last] * nums[last]
			last--
		}
	}

	return answer
}

// Runtime: 25ms, Memory Usage: 6.7mb
func sortedSquares2(nums []int) []int {
	len := len(nums)
	first, last := 0, len-1
	answer := make([]int, len)

	firstSq := nums[first] * nums[first]
	lastSq := nums[last] * nums[last]

	len--

	for first < last {
		if firstSq >= lastSq {
			answer[len] = firstSq
			first++
			firstSq = nums[first] * nums[first]
		} else {
			answer[len] = lastSq
			last--
			lastSq = nums[last] * nums[last]
		}
		len--
	}

	answer[len] = firstSq

	return answer
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
