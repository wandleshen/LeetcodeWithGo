/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max, head, tail := 0, 0, len(height)-1
	for head < tail {
		w := tail - head
		h := 0
		if height[head] < height[tail] {
			h = height[head]
			head++
		} else {
			h = height[tail]
			tail--
		}
		num := w * h
		if num > max {
			max = num
		}
	}
	return max
}

// @lc code=end

