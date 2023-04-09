/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (53.87%)
 * Likes:    679
 * Dislikes: 0
 * Total Accepted:    150.8K
 * Total Submissions: 279.8K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "1111"
 * 输出：["1.1.1.1"]
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "010010"
 * 输出：["0.10.0.10","0.100.1.0"]
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 仅由数字组成
 *
 *
 */

// @lc code=start
func restoreIpAddresses(str string) []string {
	res := []string{}
	n := len(str)
	if n > 12 || n < 4 {
		return res
	}
	//str := []byte(s)
	check := func(l, r int) bool {
		//fmt.Println(l, r)
		if l > r || r-l > 2 || (r-l > 0 && str[l] == '0') {
			return false
		}
		sum := 0
		for i := l; i <= r; i++ {
			sum = sum*10 + (int(str[i]) - '0')
			if sum > 255 {
				return false
			}
		}
		return true
	}
	var dfs func(int, int)
	dfs = func(st, pt int) {
		if pt == 3 {
			if check(st, len(str)-1) {
				//tmp := append("", str...)
				res = append(res, str)
			}
			return
		}
		for i := st; i < len(str); i++ {
			if !check(st, i) {
				continue
			}
			pt++
			tmp1 := str[:i+1]
			tmp2 := str[i+1:]
			str = tmp1 + "." + tmp2
			dfs(i+2, pt)
			str = tmp1 + tmp2
			pt--
		}
	}
	dfs(0, 0)
	return res
}

// @lc code=end

