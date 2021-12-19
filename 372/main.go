package main

import (
	"fmt"
	"math"
)

// describe: a^b % 1337
func superPow(a int, b []int) int {
	f := math.Pow(7, 10)
	return int(f) % 1337
}

//a = 2147483647, b = [2,0,0]
// 1198
// a^b % 1337 == (a % 1337) ^ b
func main() {
	fmt.Println(superPow(1, nil))
}
