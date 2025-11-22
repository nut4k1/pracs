// 58. Length of Last Word

package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	counter := 0
	flag := false
	for _, val := range s {
		if val == ' ' {
			flag = true
		} else {
			if flag {
				counter = 1
			} else {
				counter++
			}
			flag = false
		}
	}
	return counter
}

func main() {
	fmt.Println(lengthOfLastWord("me and words "))
	fmt.Println(lengthOfLastWord("nospaces"))
	fmt.Println(lengthOfLastWord("classic case"))
}
