/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	getParenthesis(n, n, "", &ans)
	return ans
}

func getParenthesis(left, right int, str string, ans *[]string) {
	if left == 0 && right == 0 {
		*ans = append(*ans, str)
		return
	}
	if left > 0 {
		getParenthesis(left-1, right, str+"(", ans)
	}
	if right > left && right > 0 {
		getParenthesis(left, right-1, str+")", ans)
	}
}

// @lc code=end

