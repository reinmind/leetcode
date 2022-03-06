package main

type point struct {
	x int
	y int
}

func setZeroes(matrix [][]int) {
	var pa []point
	for i, v := range matrix {
		for j, v2 := range v {
			if v2 == 0 {
				pa = append(pa, point{i, j})
			}
		}
	}
	for _, v := range pa {
		preset(matrix, v)
	}
}
func preset(matrix [][]int, p point) {
	x := p.x
	y := p.y
	for i := 0; i < len(matrix[x]); i++ {
		matrix[x][i] = 0
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
}

// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
func main() {

}
