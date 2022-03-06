package main

import "fmt"

func strStr(haystack string, needle string) int {
	p := -1
	if len(needle) < 1 {
		return 0
	}
	for i, j := 0, 0; i < len(haystack) && j < len(needle); {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				p = i - j
				break
			}
			continue
		}
		for j > 0 && haystack[i-j:i+1] != needle[0:j+1] {
			j--
		}
		for j == 0 && i < len(haystack) && haystack[i] != needle[0] {
			i++
		}
	}
	return p
}
func main() {
	haystack := "mississippi"
	needle := "issipi"
	// fmt.Println(haystack[1:5] == needle[0:4])
	fmt.Println(strStr(haystack, needle))
}
