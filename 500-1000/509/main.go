package main

import "fmt"

var result [31]int

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 {
		return 1
	}
	if value := result[n]; value != 0 {
		return value
	}
	tmp := fib(n-1) + fib(n-2)
	result[n] = tmp
	return tmp
}
func main() {
	fmt.Println(fib(10))
}
