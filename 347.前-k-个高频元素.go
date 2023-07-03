/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode.cn/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.51%)
 * Likes:    1558
 * Dislikes: 0
 * Total Accepted:    425.9K
 * Total Submissions: 670.6K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	max_f := 0
	n := len(nums)
	res := []int{}
	for i := 0; i < n; i++ {
		mp[nums[i]]++
		max_f = max(max_f, mp[nums[i]])
	}
	buckets := map[int][]int{}
	for k, v := range mp {
		buckets[v] = append(buckets[v], k)
	}
	for i := max_f; i > 0; i-- {
		if len(buckets[i]) > 0 {
			for k := 0; k < len(buckets[i]); k++ {
				res = append(res, buckets[i][k])
			}
			if len(res) == k {
				return res
			}
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

