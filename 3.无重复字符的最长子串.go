/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
   if s == "" {
 
      return 0
    }
    right, left, length := -1, 0, 0
    var frequency [256]int
    
    for left < len(s) {
        if right + 1 < len(s) && frequency[s[right + 1] - 'a'] == 0 {
            frequency[s[right + 1] - 'a']++
            right++
        } else {
            frequency[s[left] - 'a']--
            left++
        }
        length = max(length, right - left + 1)
    }
    return length
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
// @lc code=end

