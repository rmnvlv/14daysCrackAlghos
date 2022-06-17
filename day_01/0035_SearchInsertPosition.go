package day01

/*Easy*/

//Runtime: 4 ms, Memory Usage: 3 MB,
func searchInsert01(nums []int, target int) int {
	first := 0
	last := len(nums) - 1
	mid := first + (last-first)/2

	for first <= last {

		mid = first + (last-first)/2
		if nums[mid] < target {
			first = mid + 1
		} else if target < nums[mid] {
			last = mid - 1
		} else {
			return mid
		}

	}

	if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}

//Runtime: 4 ms, Memory Usage: 3 MB,
func searchInsert02(nums []int, target int) int {
	first := 0
	last := len(nums) - 1
	mid := first + (last-first)/2

	for first <= last {

		mid = first + (last-first)/2
		if nums[mid] < target {
			first = mid + 1
		} else if target < nums[mid] {
			last = mid - 1
		} else {
			return mid
		}

	}

	return first
}
