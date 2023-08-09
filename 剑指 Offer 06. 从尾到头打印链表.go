/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    if head == nil {
        return nil
    }
    return append(reversePrint(head.Next), head.Val)
}


 func reversePrint(head *ListNode) (res []int) {
    stack := []int{}
    for it := head; it != nil; it = it.Next {
        stack = append(stack, it.Val)
    }
    for it := len(stack)-1; it >= 0; it-- {
        res = append(res, stack[it])
    }
    return res
}