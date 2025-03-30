package main

import "fmt"

func canJump(nums []int) bool {

	//canJump := true
	jumpCount := 0

	for i := 0; i < len(nums); i++ {
		if i > jumpCount {
			return false
		}

		if jumpCount < i+nums[i] {
			jumpCount = i + nums[i]
		}
	}
	return true
}

func main() {
	case1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(case1))

	case2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(case2))

	case3 := []int{0}
	fmt.Println(canJump(case3))

	case4 := []int{2, 0, 0}
	fmt.Println(canJump(case4))
}
