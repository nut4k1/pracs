// 14. Longest Common Prefix

package main
import("fmt")

func min_length(vals []string) (min int) {
	min = len(vals[0])
	for _, val := range vals {
		if len(val) < min {
			min = len(val)
		}
	}
	return
}

func compare_rune(strs []string, idx int) bool {
	cache := strs[0][idx]
	for _, v := range strs {
		if v[idx] != cache {
			return false
		}
	}

	return true
}

func longestCommonPrefix(strs []string) (result string) {
	for i := 0; i < min_length(strs); i++ {
		if compare_rune(strs, i) {
			result += string(strs[0][i])
		} else {
			return result
		}
	}
	return
}

func main() {
	fmt.Println(longestCommonPrefix([]string { "flower", "flow", "flight" }))
	fmt.Println(longestCommonPrefix([]string { "dog", "cat", "cat" }))
	fmt.Println(longestCommonPrefix([]string { "cat", "cat", "cat" }))
}
