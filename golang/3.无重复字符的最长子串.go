/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	dic := map[byte]int{}
	n := len(s)
	l, m := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(dic, s[i-1])
		}
		for l+1 < n && dic[s[l+1]] == 0 {
			dic[s[l+1]]++
			l++
		}
		m = max(m, l-i+1)
	}
	return m
}

func max(r, length int) int {
	if r < length {
		return length
	}
	return r
}

// @lc code=end

