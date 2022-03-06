package main

import "fmt"

var cc []int
var memo []int

func minCostClimbingStairs(cost []int) int {
	cc = make([]int, 1001)
	memo = make([]int, 1001)
	// copy(cc, cost)
	for i, v := range cost {
		cc[i] = v
	}
	memo[0] = cc[0]
	memo[1] = cc[1]
	return f(len(cost))
}
func f(n int) int {
	// fmt.Println(cc[:4])
	if n < 2 {
		return cc[n]
	}
	if memo[n] != 0 {
		return memo[n]
	}
	a := f(n - 1)
	b := f(n - 2)
	if a < b {
		memo[n] = a + cc[n]
	} else {
		memo[n] = b + cc[n]
	}
	// fmt.Println(n, a, b, cc[n], memo[n])
	return memo[n]
}
func main() {
	// [10,15,20]
	// 15
	//[1,100,1,1,1,100,1,1,100,1]
	//[1,100,1,1,1,100,1,1,100,1]
	a := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(a))
}
