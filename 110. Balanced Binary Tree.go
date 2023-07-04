思路1：
平衡二叉树 -- 二叉树的每个节点的左右子树的高度差的绝对值不超过1,则二叉树是平衡二叉树。
因此,当且仅当二叉树的所有子树都是平衡二叉树时,该二叉树才是平衡二叉树,
可以使用递归的方式判断二叉树是否平衡。
先计算子树的高度,再判断左右子树的高度差是否不超过1,
然后分别递归遍历左右子节点,并判断左右子树是否平衡。
这是自顶向下的递归过程。

时间复杂度 O(n^2),空间复杂度 O(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil{
        return true
    }
    sub := height(root.Left) - height(root.Right)
    if sub >= -1 && sub <= 1 {
        return isBalanced(root.Left) && isBalanced(root.Right)
    }
    return  false
}

func height(node *TreeNode) int{
    if node == nil {
        return 0
    }
    return max(height(node.Left), height(node.Right)) + 1
}

func max(a, b int) int{
    if a >= b{
        return a
    }
    return b
}

思路2：
思路1时自顶向下递归,对于同一节点,height 函数会被重复调用,导致时间复杂度较高,
如果使用自底向上的做法,则对于每个节点,函数 height 只会被调用一次。
自底向上递归的做法类似后序遍历,对于当前遍历到的节点,
先递归地判断其左右子树是否平衡,
再判断以当前节点为根的子树是否平衡,
如果树平衡,则返回高度,否则返回-1,
若存在一棵树不平衡,则整个二叉树不平衡。

时间复杂度O(n),空间复杂度O(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    return height(root) >= 0
}

func height(node *TreeNode) int{
    if node == nil{
        return 0
    }
    left := height(node.Left)
    right := height(node.Right)
    if left == -1 || right == -1 || abs(left - right) > 1{
        return -1
    }
    return max(left, right) + 1
}

func max(a, b int) int{
    if a >= b{
        return a
    }
    return b
}

func abs(x int) int{
    if x < 0{
        return -1 * x
    }
    return x
}