package main

import "fmt"

var memo [50]int

func climbStairs(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	if n == 1 || n == 2 {
		return n
	}
	memo[n] = climbStairs(n-2) + climbStairs(n-1)
	return memo[n]
}

func main() {
	fmt.Println(climbStairs(4))
}
