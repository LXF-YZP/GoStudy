package algorithm01

// 题目描述 ：给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 题目解析:输入只有一个字符串，要求子串里面不能够有重复的元素，这里 count 都不需要定义，直接判断哈希散列里面的元素是不是在窗口内即可，是的话得移动左指针去重。

func lengthOfLongestSubstring(s string) int {

	var n = len(s)
	var ans = 0
	//创建map窗口,i为左区间，j为右区间，右边界移动
	map1 := make(map[byte]int, 0)
	j := 0
	i := 0
	for ; j < n; j++ {
		// 如果窗口中包含当前字符，
		// (map.containsKey(s.charAt(j)))
		if v, ok := map1[s[j]]; ok {
			//左边界移动到 相同字符的下一个位置和i当前位置中更靠右的位置，这样是为了防止i向左移动
			i = max(v, i)
		}
		//比对当前无重复字段长度和储存的长度，选最大值并替换
		//j-i+1是因为此时i,j索引仍处于不重复的位置，j还没有向后移动，取的[i,j]长度
		ans = max(ans, j-i+1)
		// 将当前字符为key，下一个索引为value放入map中
		// value为j+1是为了当出现重复字符时，i直接跳到上个相同字符的下一个位置，if中取值就不用+1了
		map1[s[j]] = j + 1
	}
	return ans
}

// // 哈希集合，记录每个字符是否出现过
// m := map[byte]int{}
// n := len(s)
// // 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
// rk, ans := -1, 0
// for i := 0; i < n; i++ {
// 	if i != 0 {
// 		// 左指针向右移动一格，移除一个字符
// 		delete(m, s[i-1])
// 	}
// 	for rk + 1 < n && m[s[rk+1]] == 0 {
// 		// 不断地移动右指针
// 		m[s[rk+1]]++
// 		rk++
// 	}
// 	// 第 i 到 rk 个字符是一个极长的无重复字符子串
// 	ans = max(ans, rk - i + 1)
// }
// return ans

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 方法一:
// public int lengthOfLongestSubstring(String s) {
//     if (s == null || s.length() == 0) {
//         return 0;
//     }

//     char[] sArr = s.toCharArray();
//     int[] hash = new int[128];

//     int l = 0, result = 1;
//     for (int r = 0; r < sArr.length; ++r) {
//         hash[sArr[r]]++;

//         while (hash[sArr[r]] != 1) {
//             hash[sArr[l]]--;
//             l++;
//         }

//         result = Math.max(result, r - l + 1);
//     }

//     return result;
// }

// 方法二：
// public int lengthOfLongestSubstring(String s) {
//     	int n = s.length(), ans = 0;
//         //创建map窗口,i为左区间，j为右区间，右边界移动
//         Map<Character, Integer> map = new HashMap<>();
//         for (int j = 0, i = 0; j < n; j++) {
//             // 如果窗口中包含当前字符，
//             if (map.containsKey(s.charAt(j))) {
//                 //左边界移动到 相同字符的下一个位置和i当前位置中更靠右的位置，这样是为了防止i向左移动
//                 i = Math.max(map.get(s.charAt(j)), i);
//             }
//             //比对当前无重复字段长度和储存的长度，选最大值并替换
//             //j-i+1是因为此时i,j索引仍处于不重复的位置，j还没有向后移动，取的[i,j]长度
//             ans = Math.max(ans, j - i + 1);
//             // 将当前字符为key，下一个索引为value放入map中
//             // value为j+1是为了当出现重复字符时，i直接跳到上个相同字符的下一个位置，if中取值就不用+1了
//             map.put(s.charAt(j), j+1);
//         }
//         return ans;
// }

// 核心：只增大不减小的滑动窗口
// 定义两个游标start和end，遍历字符串，end游标随着遍历无重复增大
// 若出现重复字符，则两个游标都增大index+1位（窗口大小不变，start游标滑动到重复位置后一位）
// 这时候由于窗口大小不变，窗口内可能存在重复，恰好遍历从start游标开始，如果没有重复，需要判断i+1与end的关系，超过的话，即i遍历到窗口之外，end再增大
// 遍历结束，end-start即为结果。
// 代码

// func lengthOfLongestSubstring(s string) int {
//     start,end := 0,0
//     for i := 0; i < len(s); i++ {
//         index := strings.Index(s[start:i],string(s[i]))
//         if index==-1{
//             if i+1>end{
//                 end=i+1
//             }
//         }else{
//                 start+=index+1
//                 end+=index+1
//             }
//     }
//     return end-start
// }
