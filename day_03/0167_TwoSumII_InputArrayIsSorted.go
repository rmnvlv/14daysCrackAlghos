package day03

/* Medium */

func twoSum(numbers []int, target int) []int {

	first, last := 0, len(numbers)-1

	for {
		if numbers[first]+numbers[last] > target {
			last--
		} else if numbers[first]+numbers[last] < target {
			first++
		} else {
			return []int{first + 1, last + 1}
		}
	}
}
