package algorithm01

import "math"

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
// 示例：
// 输入：s = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 进阶：
// 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

//暴力法
func minSubArrayLen01(s int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i]
		c := 1
		for j := i + 1; j < len(nums) && sum < s; j++ {
			if sum < s {
				sum += nums[j]
				c++
			}
		}
		if sum >= s {
			if count == 0 {
				count = c
			} else {
				count = min(c, count)
			}

		}
	}
	if count == 0 {
		return 0
	} else {
		return count
	}

}

//滑动窗口
func minSubArrayLen(s int, nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}
