package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	squares := [][]int{{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6}}
	// 1. if squares repeat
	for _, v := range squares {
		b := flat(board, v[0], v[1])

		if repeat(b) {
			return false
		}
	}
	// 2. if row repeat
	for _, v := range board {
		if repeat(v) {
			return false
		}
	}
	// 3. if column repeat
	b := reverse(board)
	for _, v := range b {
		if repeat(v) {
			return false
		}
	}
	return true
}
func reverse(board [][]byte) [][]byte {
	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			tmp := board[i][j]
			board[i][j] = board[j][i]
			board[j][i] = tmp
		}
	}
	return board
}
func flat(board [][]byte, p int, q int) []byte {
	a := make([]byte, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[3*i+j] = board[p+i][q+j]
		}
	}
	return a
}

func repeat(a []byte) bool {
	tag := [70]bool{}
	for _, v := range a {
		if v == '.' {
			continue
		}
		if tag[v] {
			return true
		}
		tag[v] = true
	}
	return false
}
func main() {
	fmt.Println(byte('1'))
}

/**
[[".",".","4",".",".",".","6","3","."]
,[".",".",".",".",".",".",".",".","."]
,["5",".",".",".",".",".",".","9","."]
,[".",".",".","5","6",".",".",".","."]
,["4",".","3",".",".",".",".",".","1"]
,[".",".",".","7",".",".",".",".","."]
,[".",".",".","5",".",".",".",".","."]
,[".",".",".",".",".",".",".",".","."]
,[".",".",".",".",".",".",".",".","."]]
**/
