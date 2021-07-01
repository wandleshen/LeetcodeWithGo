/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	ans, diff := 0, math.MaxInt32
	if len(nums) <= 2 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if abs(sum-target) < diff {
				ans, diff = sum, abs(sum-target)
			}
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// @lc code=end

