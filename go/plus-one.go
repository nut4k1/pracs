// 66. Plus One

package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	plus_one := true

	for i := len(digits) - 1; i >= 0; i-- {
			if plus_one {
					digits[i] = digits[i] + 1
					if digits[i] == 10 {
							digits[i] = 0
					} else {
							plus_one = false
					}
			}
	}

	result := make([]int, len(digits), len(digits) + 1)

	if plus_one {
			result = append([]int { 1 }, digits...)
	} else {
			result = digits
	}

	return result
}

func main() {
	fmt.Println(plusOne([]int { 1, 2, 3 }))
	fmt.Println(plusOne([]int { 9, 9 }))
	fmt.Println(plusOne([]int { 1, 0, 9 }))
}
