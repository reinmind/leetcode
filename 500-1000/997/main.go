package main

import "fmt"

// 1. (n-1)eveny one point to A
// 2. A point to none
func findJudge(n int, trust [][]int) int {
	m := make(map[int]int)
	folks := make([]byte, n+1)
	judge := -1
	for _, v := range trust {
		m[v[1]]++
		folks[v[0]] = 1
	}
	for i := 1; i <= n; i++ {
		if m[i] == n-1 {
			if judge == -1 {
				judge = i
			} else {
				return -1
			}
		}
		if m[i] >= n {
			return -1
		}
	}

	if folks[judge] == 1 {
		return -1
	}
	return judge
}

// 输入：n = 2, trust = [[1,2]]
// 输出：2
func main() {
	// a := [][]int{{1, 2}}
	a := [][]int{{1, 3}, {2, 3}}
	fmt.Println(findJudge(3, a))
}
