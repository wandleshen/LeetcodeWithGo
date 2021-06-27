/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	var (
		thousands = []string{"", "M", "MM", "MMM"}
		hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	)

	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}

// @lc code=end

