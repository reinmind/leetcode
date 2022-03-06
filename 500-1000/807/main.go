package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	cm := [100]int{}
	rm := [100]int{}
	count := 0
	for i := 0; i < len(grid); i++ {
		cmax := 0
		rmax := 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > cmax {
				cmax = grid[i][j]
			}
			if grid[j][i] > rmax {
				rmax = grid[j][i]
			}
		}
		cm[i] = cmax
		rm[i] = rmax
	}

	for i, v := range grid {
		for j, p := range v {
			i2 := min(cm[i], rm[j])
			if i2 > p {
				count += (i2 - p)
			}
		}
	}

	return count
}
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
