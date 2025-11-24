// 69. Sqrt(x)
package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	l := 0
	r := x / 2

	for l <= r {
		middle := l + (r-l)/2
		if middle*middle == x {
			return middle
		} else {
			if middle*middle > x {
				r = middle - 1
			} else {
				l = middle + 1
			}
		}
	}

	return l - 1
}

func main() {
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(15))
}
