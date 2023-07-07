/*
思路1：
若存在从 root 节点到叶子节点的路径和为 sum，
则存在从 root 节点的子节点到叶子节点的路径和为 sum - root.Val，
该思路满足递归的特点，递归的点在于 - 当前节点值是否等于 sum - 父节点的val，
如果当前节点就是叶子节点，那么直接判断 sum 是否等于当前节点的 val 即可。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}