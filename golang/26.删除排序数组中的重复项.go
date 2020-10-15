/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	l := len(nums)
	i := 0
	for i < l-1 {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}
	newL := len(nums)
	return newL
}

// greater solution
// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	last, finder := 0, 0
// 	for last < len(nums)-1 {
// 		for nums[finder] == nums[last] {
// 			finder++
// 			if finder == len(nums) {
// 				return last + 1
// 			}
// 		}
// 		nums[last+1] = nums[finder]
// 		last++
// 	}
// 	return last + 1
// }

// @lc code=end
