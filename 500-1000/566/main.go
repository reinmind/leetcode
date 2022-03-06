package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	r0 := len(mat)
	c0 := len(mat[0])
	size := r0 * c0
	if r*c != size {
		return mat
	}
	a := flat(mat)
	reshaped := make([][]int, r)
	for i := 0; i < r; i++ {
		begin := i * c
		end := begin + c
		reshaped[i] = a[begin:end]
	}
	return reshaped
}
func flat(mat [][]int) []int {
	var a []int
	for _, v := range mat {
		a = append(a, v...)
	}
	return a
}

func main() {
	mat := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	r := 4
	c := 2
	fmt.Println(matrixReshape(mat, r, c))
}
