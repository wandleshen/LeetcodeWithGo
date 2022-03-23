/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	left, right, maxleft, maxright := 0, -1, 0, 0
	for left < len(s) {
		//先找到回文字符的最右端
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		//向外扩展
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right-left > maxright-maxleft {
			maxright, maxleft = right, left
		}

		left = (right+left)/2 + 1
		right = left
	}

	return s[maxleft : maxright+1]
}

// @lc code=end

