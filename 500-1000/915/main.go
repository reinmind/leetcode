package main

import "fmt"

/**
给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 的长度要尽可能小。
在完成这样的分组后返回 left 的 长度 。

用例可以保证存在这样的划分方法。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-array-into-disjoint-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。**/

func partitionDisjoint(nums []int) int {
	nums_len := len(nums)
	max_left := make([]int, len(nums))
	min_right := make([]int, len(nums))
	max_tmp := nums[0]
	min_tmp := nums[nums_len-1]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max_tmp {
			max_tmp = nums[i]
		}
		max_left[i] = max_tmp
	}

	for i := nums_len - 1; i >= 0; i-- {
		if nums[i] < min_tmp {
			min_tmp = nums[i]
		}
		min_right[i] = min_tmp
	}
	// fmt.Println("", max_left)
	// fmt.Println("", min_right)
	for i := 0; i < nums_len; i++ {
		if max_left[i] <= min_right[i+1] {
			return i + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 1, 1, 0, 6, 12}
	fmt.Println(partitionDisjoint(nums))
}
