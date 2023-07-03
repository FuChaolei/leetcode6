/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode.cn/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (37.11%)
 * Likes:    2191
 * Dislikes: 0
 * Total Accepted:    355.3K
 * Total Submissions: 957.1K
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = ""
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s[i] 为 '(' 或 ')'
 *
 *
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	l, r := 0, 0
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			res = max(res, l*2)
		}
		if r > l {
			r = 0
			l = 0
		}
	}
	l = 0
	r = 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			res = max(res, l*2)
		}
		if l > r {
			r = 0
			l = 0
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

