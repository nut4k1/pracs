// 67. Add Binary

package main

import (
	"fmt"
)

func biSum(a byte, b byte, flag *bool, result []byte, index int) {
	if a == b {
		if *flag {
			result[index] = '1'
		} else {
			result[index] = '0'
		}
		*flag = a == '1'
	} else {
		if *flag {
			result[index] = '0'
		} else {
			result[index] = '1'
		}
	}
}

func biStrSum(min_str string, max_str string) []byte {
	result := make([]byte, len(max_str)+1)
	plus_one := false

	for i := 0; i < len(max_str); i++ {
		if i < len(min_str) {
			biSum(
				max_str[len(max_str)-i-1],
				min_str[len(min_str)-i-1],
				&plus_one,
				result,
				len(result)-i-1,
			)
		} else {
			biSum(
				max_str[len(max_str)-i-1],
				'0',
				&plus_one,
				result,
				len(result)-i-1,
			)
		}
	}

	if plus_one {
		result[0] = '1'
		return result
	}
	return result[1:]
}

func addBinary(a string, b string) string {
	fmt.Printf("sum of %v and %v \n", a, b)
	if len(a) > len(b) {
		return string(biStrSum(b, a))
	} else {
		return string(biStrSum(a, b))
	}
}

func main() {
	fmt.Println(addBinary("1011", "1010"))
	fmt.Println(addBinary("100", "110010"))
}
