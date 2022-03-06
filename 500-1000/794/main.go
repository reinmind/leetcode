package main

import "math"

// N('X') >= N('O')
// and
// abs(N('X') - N('O') ) < 2
// and
// Link('X') & Link('O') = 0
// and 
func validTicTacToe(board []string) bool {
	countX := 0
	countO := 0
	for _,s := range board {
		ra := []rune(s)
		for _,r := range ra {
			if r == 'X' {
				countX++;
			}
			if r == 'O' {
				countO++;
			}
		}
	}
	if countX < countO {
		return false
	}
	// and
	if countX - countO > 1 || countX - countO < -1 {
		return false
	}
	// and
	linkX := link('X',board)

	if linkX == 1 {
		if countX <= countO {
			return false
		}
	}
	linkO := link('O',board)
	if linkO == 1 {
		if countO != countX {
			return false
		}	
	}
	rs := linkX & linkO
	return rs == 0
}
func link(c rune,board []string) int{
	for i:=0;i < 3; i++ {
		r0 := rune(board[i][0])
		r1 := rune(board[i][1])
		r2 := rune(board[i][2])
		if r0 == c && r1 == c && r2 == c {
			return 1
		}
		c0 := rune(board[0][i])
		c1 := rune(board[1][i])
		c2 := rune(board[2][i])
		if c0 == c && c1 == c && c2 == c{
			return 1
		}
	}
	s0 := rune(board[0][0])
	s1 := rune(board[1][1])
	s2 := rune(board[2][2])
	if s0 == c && s1 == c && s2 == c {
		 return 1
	}

	ss0 := rune(board[2][0])
	ss1 := rune(board[1][1])
	ss2 := rune(board[0][2])
	if ss0 == c && ss1 == c && ss2 == c {
		 return 1
	}
	return 0
}
func main() {
	//board = ["O  ","   ","   "]
	// false
	validTicTacToe(["O  ","   ","   "])
}
