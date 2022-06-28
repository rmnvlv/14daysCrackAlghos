package day02

/* Medium */

/*   [1 2 3 4 5 6 7]  k = 3

[7 6 5 4 3 2 1] --> [5 6 7 4 3 2 1] --> [5 6 7 1 2 3 4]
*/

// Runtime: 39 ms, Memory Usage: 8 mb
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(arr []int, first int, last int) {
	for first < last {
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}
}
