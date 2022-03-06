package main

import "fmt"

// ransomNote = "aa", magazine = "aab"
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	r := make(map[rune]int)
	for _, x := range magazine {
		m[x]++
	}

	for _, y := range ransomNote {
		r[y]++
	}
	for k, v := range r {
		// fmt.Printf("k:%v v:%v ", k, v)
		// fmt.Printf("mk:%v mv:%v\n", k, m[k])
		if v > m[k] {
			return false
		}
	}
	return true
}

func main() {
	b := canConstruct("aaaaaaaaac", "aaaaaaabc")
	fmt.Printf("%v\n", b)
}
