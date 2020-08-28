package algorithm01

import "sort"

// 给定两个数组，编写一个函数来计算它们的交集。
// 示例 1：
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 示例 2:
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]
// 说明：
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

// 由于同一个数字在两个数组中都可能出现多次，因此需要用哈希表存储每个数字出现的次数。对于一个数字，其在交集中出现的次数等于该数字在两个数组中出现次数的最小值。
// 首先遍历第一个数组，并在哈希表中记录第一个数组中的每个数字以及对应出现的次数，然后遍历第二个数组，对于第二个数组中的每个数字，如果在哈希表中存在这个数字，则将该数字添加到答案，并减少哈希表中该数字出现的次数。
// 为了降低空间复杂度，首先遍历较短的数组并在哈希表中记录每个数字以及对应出现的次数，然后遍历较长的数组得到交集。

//方法一
func intersect(nums1 []int, nums2 []int) []int {

	//对较短的切片进行hash可以减少空间
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	// map01 := make(map[int]int)
	map01 := map[int]int{}
	for _, v := range nums1 {
		map01[v]++
	}
	slic := make([]int, 0)
	for _, va := range nums2 {
		if map01[va] != 0 {
			slic = append(slic, va)
			map01[va]--
		}
	}
	return slic
}

//方法二
func intersect2(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0
	int2 := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] > nums2[index2] {
			index2++
		} else if nums1[index1] == nums2[index2] {
			int2 = append(int2, nums1[index1])
			index1++
			index2++
		} else {
			index1++
		}
	}
	return int2
}
