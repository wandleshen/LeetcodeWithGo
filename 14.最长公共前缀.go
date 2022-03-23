/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	ans, flag := make([]byte, 0, len(strs[0])), true
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) {
				return string(ans)
			}
			if strs[0][i] != strs[j][i] {
				flag = false
			}
		}
		if flag {
			ans = append(ans, strs[0][i])
		}
	}

	return string(ans)
}

// @lc code=end

