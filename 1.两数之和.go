/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i, _ := range nums {
		j := target - nums[i]
		if _, ok := mymap[j]; ok {
			return []int{mymap[j], i}
		}
		mymap[nums[i]] = i
	}
	return nil
}

// @lc code=end

