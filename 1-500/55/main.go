package main

import "fmt"

var mem []int
var last int

func canJump(nums []int) bool {
	last = len(nums) - 1
	mem = make([]int, len(nums))
	return calc(nums, 0)
}
func calc(nums []int, i int) bool {
	if mem[i] > 0 || i >= last {
		return true
	}
	if mem[i] < 0 {
		return false
	}
	for j := 1; j <= nums[i]; j++ {
		if calc(nums, i+j) {
			mem[i+j] = 1
			return true
		} else {
			mem[i+j] = -1
		}
	}
	return false
}
func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
