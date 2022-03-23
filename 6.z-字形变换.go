/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {

	if len(s) <= numRows || numRows == 1 {
		return s
	}

	ans, temp, k, isI, isInit := make([]byte, 0, len(s)), 0, 0, true, true

	for i := 0; i < numRows; {
		j := numRows - i - 1
		isAdd := false

		if isInit {
			temp = i
			ans = append(ans, s[temp])
			isInit = false
		}

		if temp+i*2 < len(s) && (isI || j == 0) && i != 0 {
			temp += i * 2
			ans = append(ans, s[temp])
			isI = false
			isAdd = true
		}

		if temp+j*2 < len(s) && (!isI || i == 0) && j != 0 {
			temp += j * 2
			ans = append(ans, s[temp])
			k++
			isI = true
			isAdd = true
		}

		if !isAdd {
			i++
			isInit = true
			isI = false
		}
	}

	return string(ans)
}

// @lc code=end

