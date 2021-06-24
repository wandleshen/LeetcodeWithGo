/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numX := x
	num := 0
	for numX > 0 {
		num = num*10 + numX%10
		numX /= 10
	}

	return num == x
}

// @lc code=end

