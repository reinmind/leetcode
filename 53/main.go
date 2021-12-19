package main

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := -999999
	min := 999999
	value := 0
	result := nums[0]
	for _, v := range nums {
		value += v
		if value < min {
			min = value
		}
		if value > max {
			max = value
			tmp := max - min
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}
