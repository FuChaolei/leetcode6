/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 *
 * https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (56.94%)
 * Likes:    911
 * Dislikes: 0
 * Total Accepted:    184.8K
 * Total Submissions: 324.4K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
 * 输出：3
 * 解释：长度最长的公共子数组是 [3,2,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 100
 *
 *
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	res := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(dp[i][j], res)
			}
		}
	}
	return res
}
func findLength(nums1 []int, nums2 []int) int {
	getl := func(addA, addB, len int) int {
		k := 0
		res := 0
		for i := 0; i < len; i++ {
			if nums1[addA+i] == nums2[addB+i] {
				k++
			} else {
				k = 0
			}
			res = max(res, k)
		}
		return res
	}
	m, n := len(nums1), len(nums2)
	res := 0
	for i := 0; i < m; i++ {
		len := min(m-i, n)
		l := getl(i, 0, len)
		res = max(l, res)
	}
	for i := 0; i < n; i++ {
		len := min(n-i, m)
		l := getl(0, i, len)
		res = max(l, res)
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

