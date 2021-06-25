/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	/* 偷鸡法一
	match, _ := regexp.MatchString("^"+p+"$", s)
	return match
	*/

	//正经法二
	if len(p) == 0 {
		return len(s) == 0
	}

	match := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || match && isMatch(s[1:], p)
	}

	return match && isMatch(s[1:], p[1:])
}

// @lc code=end

