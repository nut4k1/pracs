// 27. Remove Element

package main
import(
	"fmt"
)

func removeElement(nums []int, val int) int {
	counter := 0

	for i := 0; i < len(nums); i++ {
			if nums[i] != val {
					nums[counter] = nums[i]
					counter++
			}
	}

	return counter
}

func main() {
	array := []int { 1, 1, 2, 2, 4, 6 }
	fmt.Println(array)
	fmt.Println(removeElement(array, 2))
	fmt.Println(array)
}