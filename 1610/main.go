package main

import (
	"fmt"
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	// z := make(map[float64]bool)

	points2 := make([][]int, len(points))
	x := location[0]
	y := location[1]
	have := 0
	index := 0
	for _, v := range points {
		if v[0] == location[0] && v[1] == location[1] {
			have += 1
			continue
		}
		points2[index] = v
		index++
	}
	points3 := points2[0 : len(points)-have]
	count := 0
	angles := make([]float64, len(points3))
	ang := float64(angle) / 180.0 * math.Pi

	for k, v := range points3 {
		dx := v[0] - x
		dy := v[1] - y
		absdx := math.Abs(float64(dx))
		absdy := math.Abs(float64(dy))
		angles[k] = math.Atan(absdy / absdx)
		// 在点下 +pi 处理
		if dy < 0 && dx < 0 {
			angles[k] += math.Pi
		}
		if dy < 0 && dx > 0 {
			angles[k] = 2*math.Pi - angles[k]
		}
		if dy > 0 && dx < 0 {
			angles[k] = math.Pi - angles[k]
		}
		if dx == 0 && dy > 0 {
			angles[k] = math.Pi / 2
		}
		if dx == 0 && dy < 0 {
			angles[k] = math.Pi + math.Pi/2
		}
		if dy == 0 && dx < 0 {
			angles[k] = math.Pi
		}
		if dy == 0 && dx > 0 {
			angles[k] = 0
		}
		//fmt.Println(dx, dy, angles[k])
	}
	//fmt.Println(len(angles))
	sort.Float64s(angles)

	for p, q := 0, 0; q < 2*len(angles) && p < len(angles); {
		i := p % len(angles)
		j := q % len(angles)
		if (angles[j]-angles[i]) > ang && q < len(angles) {
			p++
			q++
			continue
		}
		if (angles[j]-angles[i]+2*math.Pi) > ang && q >= len(angles) {
			p++
			q++
			continue
		}
		tmp := q - p + 1
		if tmp > count {
			// fmt.Println(j, i)
			count = tmp
		}
		q++
	}

	//fmt.Println(ang)
	return count + have
}
func main() {
	/**
		[[1,1],[2,2],[3,3],[4,4],[1,2],[2,1]]
	0
	[1,1]*/
	// points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
	// 3
	// fmt.Println()
	/**
		[[2,1],[2,2],[3,4],[1,1]]
	90
	[1,1]
		**/
	// points := [][]int{{0, 0}, {0, 2}}
	// points := [][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}}
	/**
			[[60,61],[58,47],[17,26],[87,97],[63,63],[26,50],[70,21],[3,89],
			[51,24],[55,14],[6,51],[64,21],[66,33],
			[54,35],[87,38],[30,0],[37,92],[92,12],
			[60,73],[75,98],[1,11],[88,24],[82,92]]
		44
		[15,50]

		[[8,67],[6,72],[7,23],[62,82],[16,54],[6,63],[55,4],[40,3],[72,12],[2,69],[61,58],[55,40],[40,55],[63,19],[45,91],[20,89],[62,72],[68,40],[75,97],[73,75],[93,38],[69,38],[7,40],[36,95],[29,88],[45,51],[19,21]]
	28
	[17,3]
			**/
	points := [][]int{{1, 1}, {2, 2}, {3, 4}, {1, 1}}
	angle := 0
	location := []int{1, 1}
	i := visiblePoints(points, angle, location)
	fmt.Println(i)
}
