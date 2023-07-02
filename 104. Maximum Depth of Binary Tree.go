思路1：
若左子树和右子树的最大深度为 l 和 r，则该二叉树的最大深度为 max(l,r)+1
而左子树和右子树的最大深度又可以以同样的方式计算
因此可以用“深度优先搜索”的方式来计算二叉树最大深度
先递归计算出左子树和右子树的最大深度，然后计算二叉树最大深度
递归在访问到空节点时退出

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

思路2：
广度优先搜索，在队列里存放当前层的所有节点，每次拓展下一层的时候，
不同于广度优先搜索的每次只从队列里拿出一个节点，
而是将队列里的所有节点都拿出来进行拓展，
这样能保证每次拓展完的时候队列里存放的是当前层的所有节点，即一层一层地进行拓展，
并用一个变量来维护拓展的次数，该二叉树的最大深度即为该变量的值。

 func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    
    dep := 0
    for len(queue) > 0{
        size := len(queue)
        for size > 0{
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil{
                queue = append(queue, node.Left)
            }
            if node.Right != nil{
                queue = append(queue, node.Right)
            }
            size--
        }
        dep++
    }
    return dep
}


思路3:
前序遍历
class Solution {
    public int maxDepth(TreeNode root) {
        return traversal(root,0);
    }

    // 前序遍历
    public int traversal(TreeNode node, int deep){
        if (node == null) return deep;
        // 每次进来就深度就 + 1
        deep++;
        int left = traversal(node.left, deep);
        int right = traversal(node.right, deep);
        return left > right ? left : right;
    }
}