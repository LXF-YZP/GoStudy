package algorithm01

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
//

// 提示：

// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) <= 1 || k <= 1 {
		return nums
	}
	q := &Queue{}
	var res []int
	for i := 0; i < len(nums); i++ {
		// 如果超过窗口，则删除
		if !q.Empty() && i >= q.Front()+k {
			q.DelFront()
		}
		for !q.Empty() && nums[i] >= nums[q.Tail()] {
			q.DelTail()
		}
		q.Insert(i)

		if i+1 >= k {
			res = append(res, nums[q.Front()])
		}
	}
	return res
}

// type Stack struct {
// 	a []int
// }

type Queue struct {
	arr []int
}

func (q *Queue) Insert(data int) {
	q.arr = append(q.arr, data)
}

func (q *Queue) DelFront() {
	if len(q.arr) > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = make([]int, 0)
	}
}

func (q *Queue) DelTail() {
	if len(q.arr) > 1 {
		q.arr = q.arr[:len(q.arr)-1]
	} else {
		q.arr = make([]int, 0)
	}
}

func (q *Queue) Front() int {
	if len(q.arr) != 0 {
		return q.arr[0]
	}
	return -1
}

func (q *Queue) Tail() int {
	if len(q.arr) != 0 {
		return q.arr[len(q.arr)-1]
	}
	return -1
}

func (q *Queue) Empty() bool {
	if len(q.arr) != 0 {
		return false
	}
	return true
}

// if nums == nil {
// 	return []int{}
// }

// window, res := []int{}, []int{}
// for i, x := range nums {

// 	if i >= k && window[0] <= i-k {
// 		window = window[1:]
// 	}
// 	for len(window) > 0 && nums[window[len(window)-1]] <= x {
// 		window = window[:len(window)-1]
// 	}
// 	window = append(window, i)
// 	if i >= k-1 {
// 		res = append(res, nums[window[0]])
// 	}
// }
// return res

// if len(nums) <= 0 {
// 	return nil
// }
// var (
// 	windows []int
// 	res     []int
// )
// //遍历数组
// for i, v := range nums {
// 	//i-k = 当窗口满时的最小下标
// 	//i-k >= 窗口最左边的数，则代表窗口已经满了
// 	if i >= k && windows[0] <= i-k {
// 		windows = windows[1:]
// 	}

// 	for {
// 		//如果窗口最后一个小于v，则去掉。为了维护了每次窗口最左边值最大
// 		if len(windows) > 0 && nums[windows[len(windows)-1]] <= v {
// 			windows = windows[:len(windows)-1]
// 		} else {
// 			break
// 		}
// 	}

// 	windows = append(windows, i)

// 	if i >= k-1 {
// 		res = append(res, nums[windows[0]])
// 	}
// }

// return res
