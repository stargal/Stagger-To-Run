package main

// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
// medium

// 这道题代码不难，难在对于题目的理解，而题目本身就比较有歧义。
// 简单来说，是找 所有 最深叶子结点的最近公共祖先。题目掉了个所有，并且第一个example的输出也有问题。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	left, right := depth(root.Left), depth(root.Right)
	if left == right {
		return root
	}
	if left > right {
		return lcaDeepestLeaves(root.Left)
	}
	return lcaDeepestLeaves(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
