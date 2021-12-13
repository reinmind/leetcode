package main

import "fmt"

func maxProfit(prices []int) int {
	var min int = 0
	var maxProf int = 0
	for i, j := range prices {
		if j-prices[min] > maxProf {
			maxProf = j - prices[min]
		}
		if j < prices[min] {
			min = i
		}
	}
	return maxProf
}
func main() {
	a := []int{7, 6, 4, 3, 1}
	i := maxProfit(a)
	fmt.Println(i)
}
