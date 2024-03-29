/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return root
    }
    invertTree(root.Left)
    invertTree(root.Right)
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
    return root
}