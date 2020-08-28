package algorithm01

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

func twoSum(nums []int, target int) []int {

	//方法一 暴力解法
	// for i, v := range nums {
	// 	for k := i + 1; k < len(nums); k++ {
	// 		if target-v == nums[k] {
	// 			return []int{i, k}
	// 		}
	// 	}
	// }
	// return []int{}

	//方法二  hash解法
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result

}
