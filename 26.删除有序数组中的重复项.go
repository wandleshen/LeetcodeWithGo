/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tail, finder := 0, 0
	for tail < len(nums)-1 {
		for nums[finder] == nums[tail] {
			finder++
			if finder == len(nums) {
				return tail + 1
			}
		}
		nums[tail+1] = nums[finder]
		tail++
	}
	return tail + 1
}

// @lc code=end

