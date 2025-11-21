// 1. Two Sum

package main
import("fmt")

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	value_to_index := make(map[int]int)

	for i, value := range nums {
		second_index, present := value_to_index[target - value]

		if present {
				result[0] = i
				result[1] = second_index
		} else {
				value_to_index[value] = i
		}
	}

	return result
}

func main() {
	fmt.Println(twoSum([]int { 1, 2, 3, 4, 5 }, 6))
	fmt.Println(twoSum([]int { 3, 2, 4 }, 6))
	fmt.Println(twoSum([]int { 3, 3 }, 6))
}
