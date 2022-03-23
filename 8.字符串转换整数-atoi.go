/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	flag, isBlankAllowed, isSignAllowed := 1, true, true
	var num int64
	num = 0

	for _, c := range s {
		if c == ' ' && isBlankAllowed {
			continue
		}

		if isSignAllowed {
			isBlankAllowed = false
			isSignAllowed = false
			if c == '-' {
				flag = -1
				continue
			} else if c == '+' {
				continue
			}
		}

		if c < '0' || c > '9' {
			break
		}

		isBlankAllowed = false
		isSignAllowed = false
		num = num*10 + int64(c-48)
		if num*int64(flag) < -(1 << 31) {
			return -(1 << 31)
		}
		if num*int64(flag) > 1<<31-1 {
			return 1<<31 - 1
		}
	}

	return int(num) * flag
}

// @lc code=end

