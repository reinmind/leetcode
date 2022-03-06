package main

/**
输入：nums = [100,4,200,1,3,2]
输出：4
*/
import (
	"fmt"
)

// probe position: left right mid
// maxLen := m[left] + m[right] + 1
// set left right the max len
// longest
func longestConsecutive(nums []int) int {

	maxLen := 0
	m := make(map[int]int)
	for _, j := range nums {
		if m[j] != 0 {
			continue
		}
		m[j] = m[j-1] + m[j+1] + 1
		//fmt.Printf("len:%v,lp:%v,rp:%v\n", m[j], j-m[j-1], j+m[j+1])
		m[j-m[j-1]] = m[j]
		m[j+m[j+1]] = m[j]
		if m[j] > maxLen {
			maxLen = m[j]
		}
	}
	return maxLen
}

func main() {
	a := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(a))
}
