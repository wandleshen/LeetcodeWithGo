/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	var r = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	num, temp, ans := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-i-1 : len(s)-i]
		num = r[char]
		if num >= temp {
			ans += num
		} else {
			ans -= num
		}
		temp = num
	}
	return ans

}

// @lc code=end

