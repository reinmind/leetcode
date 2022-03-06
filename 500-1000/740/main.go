package main

import "fmt"

var count map[int]int
var max int
var mem map[int]int

func deleteAndEarn(nums []int) int {
	count = make(map[int]int)
	max = 0
	mem = make(map[int]int)
	for _, v := range nums {
		if v > max {
			max = v
		}
		count[v]++
	}
	dp := make([]int, max+1)
	dp[0] = count[0]
	if count[1] > count[0] {
		dp[1] = count[1]
	} else {
		dp[1] = count[0]
	}
	for i := 2; i <= max; i++ {
		if count[i] == 0 {
			continue
		}
		a := dp[i-2] + count[i]
		b := dp[i-1]
		if a > b {
			dp[i] = a
		} else {
			dp[i] = b
		}
	}
	return dp[max]
}
func calc(i int) int {
	if i > max {
		return mem[i]
	}
	if count[i] == 0 {
		return calc(i + 1)
	}
	var a, b int
	if mem[i+2] > 0 {
		a = count[i]*i + mem[i+2]
	} else {
		a = count[i]*i + calc(i+2)
	}
	if mem[i+1] > 0 {
		b = mem[i+1]
	} else {
		b = calc(i + 1)
	}
	if a > b {
		mem[i] = a
	} else {
		mem[i] = b
	}
	return mem[i]
}

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	i := deleteAndEarn(nums)
	fmt.Println(i)
}
