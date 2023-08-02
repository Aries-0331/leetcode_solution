/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var depth func(*TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left  := depth(node.Left)
        right := depth(node.Right)
        res = max(res, left+right)
        return max(left, right) + 1
    }
    depth(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}