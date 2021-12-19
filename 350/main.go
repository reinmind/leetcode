package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	var a []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			a = append(a, nums1[i])
			i++
			j++
			continue
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return a
}
func main() {
	nums1 := []int{3, 2, 2, 1}
	nums2 := []int{2, 2, 2}
	fmt.Println(intersect(nums1, nums2))
}
