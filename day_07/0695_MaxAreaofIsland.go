package day07

func maxAreaOfIsland(grid [][]int) int {

	lenCol := len(grid)
	lenRow := len(grid[0])
	count := 0

	var counter func(row, col int)

	counter = func(row, col int) {
		count += 1
		grid[row][col] = 0

		if row != 0 && grid[row-1][col] == 1 { // Up
			counter(row-1, col)
		}

		if col != 0 && grid[row][col-1] == 1 { // Left
			counter(row, col-1)
		}

		if row != lenCol+1 && grid[row+1][col] == 1 { // Down
			counter(row+1, col)
		}

		if col != lenRow+1 && grid[row][col+1] == 1 { // Right
			counter(row, col+1)
		}
	}

	max := 0

	for i := 0; i < lenCol; i++ {
		for j := 0; j < lenRow; j++ {
			if grid[i][j] == 1 {
				count = 0
				counter(i, j)
				if count > max {
					max = count
				}
			}
		}
	}

	return max

}
