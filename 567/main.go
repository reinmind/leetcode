package main

import "fmt"

func main() {
	s1 := "abcd"
	s2 := "adcabab"
	b := checkInclusion(s1, s2)
	fmt.Println(b)
}
func equals(m1 map[int]int, m2 map[int]int) bool {
	for i := 0; i < 26; i++ {
		if m1[i] != m2[i] {
			return false
		}
	}
	return true
}
func checkInclusion(s1 string, s2 string) bool {
	ls1 := len(s1)
	ls2 := len(s2)
	if ls1 > ls2 {
		return false
	}
	m1 := make(map[int]int)
	for _, v := range s1 {
		m1[int(v-'a')]++
	}

	m2 := make(map[int]int)
	// init m2
	for i := 0; i < ls1; i++ {
		m2[int(s2[i]-'a')]++
	}
	for i, j := 0, ls1-1; ; {
		if equals(m1, m2) {
			return true
		}
		if j == ls2-1 {
			return false
		}
		m2[int(s2[i]-'a')]--
		i++
		j++
		m2[int(s2[j]-'a')]++
	}
}
