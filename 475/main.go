package main

import (
	"fmt"
	"math"
	"sort"
)

// 每个房子找到最近的加热器，即为最小半径
// 0: 什么都没
// 1: 有房子
// 2: 有加热器
// 3: 有房子和加热器
func findRadius(houses []int, heaters []int) int {
	minRadius := math.MaxInt32
	sort.Ints(houses)
	sort.Ints(heaters)
	point := make(map[int]byte)
	mixed := make([]int, 60000)
	index := 0
	for i, j := 0, 0; i < len(houses) || j < len(heaters); index++ {
		if j == len(heaters) && i < len(houses) || houses[i] > heaters[j] {
			mixed[index] = houses[i]
			point[mixed[index]] += 1
			i++
			continue
		}
		if i == len(houses) && j < len(heaters) || houses[i] <= heaters[j] {
			mixed[index] = heaters[j]
			point[mixed[index]] += 2
			j++
			continue
		}
	}

	// findHeaterRadius(house int)
	for k, v := range mixed {
		// has house
		if point[v]%2 == 1 {

		}
	}
	return 1
}
func findHeaterRadius(house int) {

}

// 输入: houses = [1,2,3,4], heaters = [1,4]
// 输出: 1
func main() {
	houses := []int{1, 2, 3, 4}
	heaters := []int{1, 4}
	fmt.Println(findRadius(houses, heaters))
}
