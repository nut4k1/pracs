// 35. Search Insert Position

package main

import (
	"fmt"
)

func recSearch(nums []int, target int, shift int) (result int) {
	fmt.Println(nums)
	if len(nums) == 1 {
		if nums[0] < target {
			result = shift + 1
		} else {
			result = shift
		}
		return
	}

	middle := (len(nums)) / 2
	if nums[middle] > target {
		result = recSearch(nums[:middle], target, shift)
	} else {
		result = recSearch(nums[middle:], target, shift + middle)
	}
	return
}

func searchInsert(nums []int, target int) int {
	return recSearch(nums, target, 0)
}

func main() {
	array := []int { 1, 3, 5, 6 }
	fmt.Println(searchInsert(array, 4))
	fmt.Println(searchInsert(array, 5))
}
