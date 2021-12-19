package main

import "math"

func containsDuplicate(nums []int) bool {
	a := [math.MaxInt]bool{}
	for _, v := range nums {
		if !a[v] {
			a[v] = true
		} else {
			return true
		}
	}
	return false
}
