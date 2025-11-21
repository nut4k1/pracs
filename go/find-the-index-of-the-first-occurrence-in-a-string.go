// 28. Find the Index of the First Occurrence in a String

package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	result := -1

	j := 0
	for i := 0; i < len(haystack); i++ {
		if j == len(needle) {
			return result
		}

		if haystack[i] == needle[j] {
			if j == 0 {
				result = i
			}
			j++
		} else {
			if j != 0 {
				i = result
			}

			result = -1
			j = 0
		}
	}

	if j == len(needle) {
		return result
	}

	return -1
}

func main() {
	haystack := "leetcode"
	needle := "leeto"
	fmt.Println(strStr(haystack, needle))

	haystack = "sadbutsad"
	needle = "sad"
	fmt.Println(strStr(haystack, needle))

	haystack = "mississippi"
	needle = "issip"
	fmt.Println(strStr(haystack, needle))

	haystack = "a"
	needle = "a"
	fmt.Println(strStr(haystack, needle))
}
