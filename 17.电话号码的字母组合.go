/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var (
	ans     = []string{}
	letters = []string{
		"",     //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	ans = []string{}
	findComb(&digits, 0, "")
	return ans
}

func findComb(digits *string, index int, subst string) {
	if index == len(*digits) {
		ans = append(ans, subst)
		return
	}
	num := (*digits)[index]
	letter := letters[num-'0']
	for i := 0; i < len(letter); i++ {
		findComb(digits, index+1, subst+string(letter[i]))
	}
	return
}

// @lc code=end

