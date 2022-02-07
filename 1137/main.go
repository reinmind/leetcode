package main

import "fmt"

var memo [40]int

func tribonacci(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	result := tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	memo[n] = result
	return result
}
func main() {
	fmt.Println(tribonacci(25))
}
