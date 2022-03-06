package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	tag := (dividend * divisor) < 0
	divisor = int(math.Abs(float64(divisor)))
	dividend = int(math.Abs(float64(dividend)))
	divisor_ori := divisor
	n := 0
	for ; dividend > divisor; n++ {
		divisor = divisor << 1
	}
	//fmt.Printf("divisor: %v\n", divisor)
	m := 0
	for ; divisor > dividend; m++ {
		divisor -= divisor_ori
	}
	//fmt.Printf("n:%v m: %v\n", n, m)
	result := 1<<n - m
	if tag {

		result = -result
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}
func main() {
	i := divide(7, 3)
	fmt.Printf("%v\n", i)
}
