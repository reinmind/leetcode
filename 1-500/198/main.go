// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
package main

import "fmt"

var max []int

func rob(nums []int) int {
	max = make([]int, 500)
	for k := range max {
		max[k] = -1
	}
	return calc(nums, 0)
}

// 根据当前位置判断最大值
func calc(nums []int, i int) int {
	if i >= len(nums) {
		return 0
	}
	if max[i] >= 0 {
		return max[i]
	}
	a := calc(nums, i+1)
	b := calc(nums, i+2) + nums[i]
	if a > b {
		max[i] = a
	} else {
		max[i] = b
	}
	return max[i]
}

func main() {
	nums := []int{1, 2, 3, 1}
	i := rob(nums)
	fmt.Println(i)
	// fmt.Println(max)
}
