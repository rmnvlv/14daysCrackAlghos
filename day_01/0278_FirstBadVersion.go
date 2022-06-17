package day01

/*Easy*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

//Runtime: 0 ms, Memory Usage: 1.9 MB
func firstBadVersion(n int) int {
	first := 1
	last := n

	for first < last {
		midle := first + (last-first)/2

		if isBadVersion(midle) {
			last = midle
		} else {
			first = midle + 1
		}
	}

	return first
}
