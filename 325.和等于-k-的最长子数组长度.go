/*
 * @lc app=leetcode.cn id=325 lang=golang
 *
 * [325] 和等于 k 的最长子数组长度
 *
 * https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (51.89%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    16.3K
 * Total Submissions: 31.4K
 * Testcase Example:  '[1,-1,5,-2,3]\n3'
 *
 * 给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,-1,5,-2,3], k = 3
 * 输出: 4
 * 解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [-2,-1,2,1], k = 1
 * 输出: 2
 * 解释: 子数组 [-1, 2] 和等于 1，且长度最长。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^5
 * -10^4 <= nums[i] <= 10^4
 * -10^9 <= k <= 10^9
 *
 *
 */

// @lc code=start
func maxSubArrayLen(nums []int, k int) int {
	res := 0
	n := len(nums)
	dp := map[int]int{}
	dp[0] = 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if _, ok := dp[sum]; !ok {
// 不用i的反例
// [-2,-1,2,1]
// 1
// Answer
// 3
// Expected Answer
// 2
			dp[sum] = i
		}
		if _, ok := dp[sum-k]; ok {
			res = max(res, i-dp[sum-k]+1)
		}
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

