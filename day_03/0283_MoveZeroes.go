package day03

/* Easy */

//Runtime: 35 ms, Memory Usage: 7 MB
func moveZeroes(nums []int) {
	index := 0
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

//Runtime: 26 ms, Memory Usage: 6.6 MB
func moveZeroes2(nums []int) {
	index := 0
	for _, num := range nums {
		if num != 0 {
			nums[index] = num
			index++
		}
	}

	for index < len(nums) {
		nums[index] = 0
		index++
	}
}

// Runtime: 27 ms Memory Usage: 7.2 MB
func moveZeroes3(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
}

//Runtime: 27 ms, Memory Usage: 6.5 MB
func moveZeroes4(nums []int) {
	index := 0
	for i, num := range nums {
		if num != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
}

func moveZeroes5(nums []int) {
	len := len(nums)

	fast, slow := 0, 0

	for fast < len {
		if nums[fast] != 0 && nums[slow] == 0 {
			nums[slow] = nums[fast]
			nums[fast] = 0
		}

		if nums[slow] != 0 {
			slow++
		}

		fast++
	}
}
