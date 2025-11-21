// 26. Remove Duplicates from Sorted Array

package main
import(
	"fmt"
)

func removeDuplicates(nums []int) int {
	counter := 0
	cache := map[int]int{}

	for i := 0; i < len(nums); i++ {
			if _, ok := cache[nums[i]]; !ok {
					cache[nums[i]] = i
					nums[counter] = nums[i]
					counter++
			}
	}

	return counter
}

func main() {
	array := []int { 1, 1, 2, 2, 4, 6 }
	fmt.Println(removeDuplicates(array))
	fmt.Println(array)
}