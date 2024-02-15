package main

import "fmt"

var rs [100][100]int

func removeBoxes(boxes []int) int {
	rs = [100][100]int{}
	return maxSquare(boxes, 0, len(boxes)-1)
}
func maxSquare(boxes []int, p int, q int) int {
	if q < 0 || p == len(boxes) {
		return 0
	}
	tag := true
	if rs[p][q] != 0 {
		return rs[p][q]
	}
	for i := p + 1; i <= q; i++ {
		if boxes[i] != boxes[p] {
			tag = false
			break
		}
	}
	if tag {
		t := q - p + 1
		rs[p][q] = t * t
		return rs[p][q]
	}
	max := 0
	for i := p; i <= q; i++ {
		for j := i; j <= q; j++ {
			tmp := maxSquare(boxes, p, i-1) + maxSquare(boxes, i, j) + maxSquare(boxes, j+1, q)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max * max
}
func main() {
	a := []int{1, 3, 2, 2, 2, 3, 4, 3, 1}
	// result 9
	i := removeBoxes(a)
	fmt.Println(i)
}
