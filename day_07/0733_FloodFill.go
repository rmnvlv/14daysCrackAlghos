package day07

// Runtime: 3 ms, Memory Usage: 4.1 MB Recursive DFS
func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	lenCol := len(image)
	lenRow := len(image[0])

	if image[sr][sc] == color {
		return image
	}

	mainColour := image[sr][sc]

	var fillCels func(sr, sc int)

	fillCels = func(sr, sc int) {

		image[sr][sc] = color // center

		if sr != 0 && image[sr-1][sc] == mainColour { // Up
			fillCels(sr-1, sc)
		}

		if sc != 0 && image[sr][sc-1] == mainColour { // Left
			fillCels(sr, sc-1)
		}

		if sr != lenCol-1 && image[sr+1][sc] == mainColour { //Down
			fillCels(sr+1, sc)
		}

		if sc != lenRow-1 && image[sr][sc+1] == mainColour { // Right
			fillCels(sr, sc+1)
		}
	}

	fillCels(sr, sc)
	return image
}
