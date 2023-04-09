/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode.cn/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (58.19%)
 * Likes:    2077
 * Dislikes: 0
 * Total Accepted:    670.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isSymmetric(root *TreeNode) bool {
// 	var check func(*TreeNode, *TreeNode) bool
// 	check = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		}
// 		if left == nil || right == nil || right.Val != left.Val {
// 			return false
// 		}
// 		return check(left.Left, right.Right) && check(left.Right, right.Left)
// 	}
// 	return check(root, root)
// }
func isSymmetric(root *TreeNode) bool {
	st := []*TreeNode{}
	st = append(st, root)
	st = append(st, root)
	for len(st) > 0 {
		tmp1 := st[0]
		tmp2 := st[1]
		st = st[2:]
		if tmp1 == nil && tmp2 == nil {
			continue
		}
		if tmp1 == nil || tmp2 == nil || tmp1.Val != tmp2.Val {
			return false
		}
		st = append(st, tmp1.Left)
		st = append(st, tmp2.Right)
		st = append(st, tmp1.Right)
		st = append(st, tmp2.Left)
	}
	return true
}

// @lc code=end

