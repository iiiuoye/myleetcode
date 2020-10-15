/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

// func twoSum(nums []int, target int) []int {
// 	v := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		dif := target - nums[i]
// 		c, ok := v[dif]
// 		if ok == true {
// 			return []int{c, i}
// 		}
// 		v[nums[i]] = i
// 	}
// 	return []int{-1, -1}
// }
package golang
func twoSum(nums []int, target int) []int {
	l, r := 0, 1
	for l < len(nums) {
		r = l + 1
		for r < len(nums) {
			if nums[l]+nums[r] == target {
				return []int{l, r}
			}
			r++
		}
		l++
		if{

		}else{
			
		}
	}
	return []int{-1, -1}
}

// @lc code=end
