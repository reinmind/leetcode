/**
*你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。



示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums =
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [1,2,3]
输出：3
*/
package main

import "fmt"

var end int
var max []int

func rob(nums []int) int {
	max = make([]int, 110)
	for k := range max {
		max[k] = -1
	}
	len := len(nums)
	// chose 0
	end = len - 2
	a := calc(nums, 2) + nums[0]
	// not chose 0
	end = len - 1
	for k := range max {
		max[k] = -1
	}

	b := calc(nums, 1)
	if a > b {
		return a
	}
	return b
}

func calc(nums []int, i int) int {
	if max[i] > -1 {
		return max[i]
	}
	if i > end {
		return 0
	}
	// chose i
	a := nums[i] + calc(nums, i+2)
	// not chose i
	b := calc(nums, i+1)
	if a > b {
		max[i] = a
	} else {
		max[i] = b
	}
	return max[i]
}

func main() {
	nums := []int{2, 3, 2}
	i := rob(nums)
	fmt.Println(i)
}
