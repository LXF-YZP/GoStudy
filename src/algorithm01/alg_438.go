package algorithm01

// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
// 说明：
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
// 示例 1:
// 输入:
// s: "cbaebabacd" p: "abc"
// 输出:
// [0, 6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
// 示例 2:
// 输入:
// s: "abab" p: "ab"
// 输出:
// [0, 1, 2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。

func findAnagrams(s string, p string) []int {

	// 输入参数有效性判断
	if len(s) < len(p) {
		return []int{}
	}

	// 申请一个散列，用于记录窗口中具体元素的个数情况
	// 这里用数组的形式呈现，也可以考虑其他数据结构
	// char[] sArr = s.toCharArray();
	// char[] pArr = p.toCharArray();

	hash := make([]int, 26)
	// int[] hash = new int[26];

	for i := 0; i < len(p); i++ {
		hash[p[i]-'a']++
	}

	// l 表示左指针
	// count 记录当前的条件，具体根据题目要求来定义
	// result 用来存放结果

	// List<Integer> results = new ArrayList<>();
	results := make([]int, 0)
	l := 0
	count := 0
	pLength := len(p)
	for r := 0; r < len(s); r++ {
		// 更新新元素在散列中的数量
		hash[s[r]-'a']--
		// 根据窗口的变更结果来改变条件值
		if hash[s[r]-'a'] >= 0 {
			count++
		}
		// 如果当前条件不满足，移动左指针直至条件满足为止
		if r > pLength-1 {
			hash[s[l]-'a']++
			if hash[s[l]-'a'] > 0 {
				count--
			}
			l++
		}
		// 更新结果
		if count == pLength {
			// results.add(l)
			results = append(results, l)
		}
	}
	return results

}

// // 输入参数有效性判断
// if (s.length() < p.length()) {
// 	return new ArrayList<Integer>();
// }

// // 申请一个散列，用于记录窗口中具体元素的个数情况
// // 这里用数组的形式呈现，也可以考虑其他数据结构
// char[] sArr = s.toCharArray();
// char[] pArr = p.toCharArray();

// int[] hash = new int[26];

// for (int i = 0; i < pArr.length; ++i) {
// 	hash[pArr[i] - 'a']++;
// }

// // l 表示左指针
// // count 记录当前的条件，具体根据题目要求来定义
// // result 用来存放结果
// List<Integer> results = new ArrayList<>();
// int l = 0, count = 0, pLength = p.length();
// for (int r = 0; r < sArr.length; ++r) {
// 	// 更新新元素在散列中的数量
// 	hash[sArr[r] - 'a']--;
// 	// 根据窗口的变更结果来改变条件值
// 	if (hash[sArr[r] - 'a'] >= 0) {
// 		count++;
// 	}
// 	// 如果当前条件不满足，移动左指针直至条件满足为止
// 	if (r > pLength - 1) {
// 		hash[sArr[l] - 'a']++;
// 		if (hash[sArr[l] - 'a'] > 0) {
// 			count--;
// 		}
// 		l++;
// 	}
// 	// 更新结果
// 	if (count == pLength) {
// 		results.add(l);
// 	}
// }
// return results;
