思路：
题目给定的数组是按照升序排列的有序数组，因此可以判断数组是二叉搜索树的中序遍历序列，
可以选择中间数字作为二叉搜索树的根节点，这样分给左右子树的数字个数相同或只相差1，可以使树保持平衡，
如果数组长度是奇数，则根节点的选择是唯一的，
如果数组长度是偶数，则可以选择中间位置左边的数字作为根节点，或选择中间位置右边的数字作为根节点，
选择不同的数字作为根节点创建的平衡二叉搜索树是不同的，但均符合题意。
确定平衡二叉树的根节点后其余数字分别位于平衡二叉树的左子树和右子树中，
左子树和右子树分别也是平衡二叉树，因此可以通过递归的方式创建平衡二叉搜索树。
递归的基准情况是平衡二叉搜索树不包含任何数字，此时平衡二叉搜索树为空。
在给定中序遍历序列数组的情况下，每个子树中的数字在数组中一定是连续的，
因此可以通过数组下标范围确定子树包含的数字，下标范围为[left, right],
对于整个中序遍历序列，下标范围从 left=0 到 right=len(nums)-1，当 left>right 时，平衡二叉搜索树为空。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedArrayToBST(nums []int) *TreeNode {
    return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, left, right int) *TreeNode{
    if left > right{
        return nil
    }
    mid := (left + right)/2
    root := &TreeNode{Val: nums[mid]}
    root.Left = bst(nums, left, mid-1)
    root.Right = bst(nums, mid+1, right)
    return root
}