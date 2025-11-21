// 13. Roman to Integer

package main
import("fmt")

var int_by_romans = map[string]int {
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0

	for i, v := range s {
		result += int_by_romans[string(v)]
		if i == 0 { continue }
		if int_by_romans[string(s[i - 1])] < int_by_romans[string(v)] {
		result -= int_by_romans[string(s[i - 1])] * 2
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("XX"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XIX"))
}
