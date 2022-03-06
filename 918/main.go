package main

import "fmt"

func main() {
	nums := []int{1, -3, 6, -2, -3}
	i := maxSubarraySumCircular(nums)
	fmt.Println(i)
}
func maxSubarraySumCircular(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	if max < 0 {
		return max
	}
	minSum := nums[0]
	maxSum := nums[0]
	sum := 0
	result := nums[0]
	for _, v := range nums {
		sum += v
		if sum < minSum {
			minSum = sum
		}
		tmp := sum - minSum
		if tmp > result {
			if tmp > result {
				result = tmp
				fmt.Println(sum, maxSum, minSum)
			}
		}

	}
	sum += nums[0]
	if sum < minSum {
		minSum = sum
	}
	tmp := sum - minSum
	if tmp > result {
		if tmp > result {
			result = tmp
			fmt.Println(sum, maxSum, minSum)
		}
	}
	return result
}
