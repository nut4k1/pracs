// 9. Palindrome Number

package main
import("fmt")

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	reversed := 0

	for num != 0 {
		reversed = 10*reversed + num%10
		num = num / 10
	}

	return (x == reversed)
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10301))
}
