package day01

/* Easy */

//Runtime: 38 ms, Memory Usage: 6.7 MB
func search01(nums []int, target int) int {
	first := 0
	last := len(nums) - 1

	for first <= last {
		mid := first + (last-first)/2

		if nums[mid] < target {
			first = mid + 1
		} else if nums[mid] > target {
			last = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

//Runtime: 29 ms, Memory Usage: 6.7 MB ?? why faster
func search02(nums []int, target int) int {
	first := 0
	last := len(nums) - 1

	for first <= last {
		mid := first + (last-first)/2

		if nums[mid] > target {
			last = mid - 1
		} else if nums[mid] < target {
			first = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
