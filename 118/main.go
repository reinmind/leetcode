package main

import "fmt"

func generate(numRows int) [][]int {
	var state []int
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		// fmt.Println(state[:])
		a := make([]int, i)
		a[0] = 1
		a[i-1] = 1
		for j := 1; j < i-1; j++ {
			a[j] = state[j-1] + state[j]
		}
		state = a
		result[i-1] = state
	}
	return result
}
func main() {
	fmt.Println(generate(1))
}
