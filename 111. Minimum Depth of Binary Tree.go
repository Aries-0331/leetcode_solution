/*
思路：
叶子节点是指没有子节点的节点，即左孩子和右孩子都为 nil 的节点
当 root 节点左右孩子都为空时，返回 1
当 root 节点左右孩子中有一个为空时，返回非空孩子节点的深度
当 root 节点左右孩子都不为空时，返回左右孩子节点深度较小的深度值
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

简化后的版本：
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftD := minDepth(root.Left)
    rightD := minDepth(root.Right)
    if root.Left == nil {
        return rightD + 1
    }
    if root.Right == nil {
        return leftD + 1
    }
    return min(leftD, rightD) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}