/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid float64 = 0
	s := append(nums1, nums2...)
	sort.Ints(s)
	l := len(s)
	if l%2 == 0 {
		mid = float64(s[l/2-1]+s[l/2]) / 2.0
	} else {
		mid = float64(s[l/2])
	}

	return mid
}

// @lc code=end

