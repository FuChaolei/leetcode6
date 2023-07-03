/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.54%)
 * Likes:    2912
 * Dislikes: 0
 * Total Accepted:    606.4K
 * Total Submissions: 782.2K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	path := ""
	var dfs func(int, int)
	dfs = func(l, r int) {
		if l == 0 && r == 0 {
			res = append(res, path)
		}
		if r < l {
			return
		}
		if l > 0 {
			path = path + "("
			dfs(l-1, r)
			path = path[:len(path)-1]
		}
		if r > 0 {
			path = path + ")"
			dfs(l, r-1)
			path = path[:len(path)-1]
		}
	}
	dfs(n, n)
	return res
}

// @lc code=end

