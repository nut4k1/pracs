// 20. Valid Parentheses

package main
import("fmt")

func isValid(s string) bool {
	if len(s) == 0 || len(s) % 2 == 1 {
		return false
	}

	pairs := map[rune]rune {
		'(': ')',
		'{': '}',
		'[': ']',
	}

	cache := []rune{}

	for _, v := range s {
		if _, ok := pairs[v]; ok {
			cache = append(cache, v)
		} else if len(cache) == 0 || pairs[cache[len(cache) - 1]] != v {
			return false
		} else {
			cache = cache[:len(cache) - 1]
		}
	}

	return len(cache) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]"))
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("(((]]"))
}
