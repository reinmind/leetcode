package main

import "fmt"

/**
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

假设你总是可以到达数组的最后一个位置。
*/
func main() {
	nums := []int{2, 0, 2, 0, 1}
	i := jump(nums)
	fmt.Println(i)
	fmt.Println(mem)
}

var mem map[int]int
var last int

func calc(nums []int, i int) int {
	if i >= last {
		return 0
	}
	if mem[i] > 0 {
		return mem[i]
	}
	min := 9999
	// 从位置i开始跳最少的次数
	// j为能够跳的步数
	for j := 1; j <= nums[i]; j++ {
		// t为跳到 i+j 位置的最少次数
		t := mem[i+j]
		if t < 1 {
			t = calc(nums, i+j)
		}
		mem[i+j] = t
		if t < min {
			min = t
		}
	}
	return min + 1
}

func jump(nums []int) int {
	mem = make(map[int]int)
	last = len(nums) - 1
	return calc(nums, 0)
}
