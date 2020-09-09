package algorithm01

//滑动窗口算法可以将嵌套的循环问题，转换为单循环问题，降低时间复杂度。
// 用来解决数组，字符串的子元素问题
// 我们在一个循环中计算出了长度为 'k' 的子数组的最大总和，它的时间复杂度是 O（n）。我们可以使用滑动窗口算法解决 查找最大 / 最小 k 子阵列，XOR，乘积，求和等问题。

// 给定一个整数数组，计算长度为 'k' 的连续子数组的最大总和。
// 输入：arr [] = {100,200,300,400}
//      k = 2
// 输出：700
// 解释：300 + 400 = 700

//暴力法
// func maxSum(arr []int, k int) int {
// 	n := len(arr)
// 	maxSum := math.MinInt32
// 	for i := 0; i < n-k+1; i++ {
// 		currentSum := 0
// 		for j := 0; j < k; j++ {
// 			// currentSum = currentSum + arr[i+j]
// 			currentSum = currentSum + arr[i+j]
// 		}
// 		maxSum = max(currentSum, maxSum) //max方法在另外的go文件中
// 	}
// 	return maxSum
// }

func maxSum01(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return -1
	}

	// 计算出第一个窗口的值
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	sum := maxSum
	for i := k; i < n; i++ {
		// 新窗口的和 = 前一个窗口的和 + 新进入窗口的值 - 移出窗口的值
		sum += arr[i] - arr[i-k]
		maxSum = max(maxSum, sum)
	}

	return maxSum
}

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
// 进阶：
// 你能在线性时间复杂度内解决此题吗？
// 示例:
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:
//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

//暴力法
// public int[] maxSlidingWindow(int[] nums, int k) {
// 	int n = nums.length;
// 	if (n * k == 0) return new int[0];

// 	int [] output = new int[n - k + 1];
// 	for (int i = 0; i < n - k + 1; i++) {
// 		int max = Integer.MIN_VALUE;
// 		for(int j = i; j < i + k; j++)
// 			max = Math.max(max, nums[j]);
// 		output[i] = max;
// 	}
// 	return output;
// }

// 给定一个字符串 S 和一个字符串 T，请在 S 中找出包含 T 所有字母的最小子串。(minimum-window-substring)
// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"
// 这个问题让我们无法按照上面求最大和的方法进行查找，因为它不是给定了窗口大小让你找对应的值，而是给定了对应的值，让你找最小的窗口。
// 我们仍然可以使用滑动窗口算法，只不过需要换一个思路。
// 既然是找最小的窗口，我们先定义一个最小的窗口，也就是长度为 0 的窗口
