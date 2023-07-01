```
 递归的难点在于找到可以递归的点，思路：
 1. 如何判断一棵树是否为对称二叉树？--如果所给根节点，为空，则对称。如果不为空，当它的左子树与右子树对称时，它对称；
 2. 如何判断左子树与右子树是否对称？--如果左子树的左孩子与右子树的右孩子对称，左子树的右孩子与右子树的左孩子对称，则左子树与右子树对称；
 3. 递归点--在尝试判断左子树与右子树对称的条件时，发现其与两树的孩子的对称情况有关系；
 4. 定义递归函数，func Fun（左子树，右子树）bool{左子树节点值等于右子树节点值，且 Fun（左子树的左孩子，右子树的右孩子），Fun（左孩子的右子树，右孩子的左子树）均为真，才返回真}；
```
 /**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isChildSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val == right.Val {
		if isChildSymmetric(left.Left, right.Right) && isChildSymmetric(left.Right, right.Left) {
			return true
		}
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
    return isChildSymmetric(root.Left, root.Right)
}


