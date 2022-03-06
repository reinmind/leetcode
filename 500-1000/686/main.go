package main

import "fmt"

func repeatedStringMatch(a string, b string) int {
	m := make(map[rune]bool)
	lena := len(a)
	lenb := len(b)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			return -1
		}
	}

	start := 0

	for ; start < lena; start++ {
		j := 0
		p := start
		for i := p; j < lenb; {
			if a[i] != b[j] {
				break
			}
			p++
			i = p % lena
			j++
		}
		if j == lenb {
			return (p-1)/lena + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("k", "kkk"))
}
