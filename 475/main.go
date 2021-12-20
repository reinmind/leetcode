package main

import (
	"fmt"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	maxRadius := -1
	for _, v := range houses {
		closetHeaterDistance := binarySearch(heaters, 0, len(heaters)-1, v)
		// fmt.Println(closetHeaterDistance)
		if closetHeaterDistance > maxRadius {
			maxRadius = closetHeaterDistance
		}
	}
	return maxRadius
}

func binarySearch(heaters []int, a int, b int, value int) int {
	gap := b - a
	if gap < 2 {
		left := value - heaters[a]
		if left < 0 {
			left = -left
		}
		right := heaters[b] - value
		if right < 0 {
			right = -right
		}
		if left < right {
			return left
		}
		return right
	}
	mid := (a + b) / 2
	if heaters[mid] > value {
		return binarySearch(heaters, a, mid, value)
	}
	return binarySearch(heaters, mid, b, value)
}

func main() {
	houses := []int{1, 999}
	heaters := []int{499, 500, 501}
	fmt.Println(findRadius(houses, heaters))
}
