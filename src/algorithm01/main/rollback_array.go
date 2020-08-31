package main

import "fmt"

func main() {

	int2 := []int{1, 2, 3}
	fmt.Println(int2)
	int3 := rollbackArray(int2)
	fmt.Println(int3)

}

func rollbackArray(nums []int) []int {
	var i, j = 0, len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
