/**
思路1：双指针
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    var cur *ListNode
    pre := head
    for pre != nil {
        tmp := pre.Next
        pre.Next = cur
        cur = pre
        pre = tmp
    }
    return cur
}

/**
思路2：递归
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    tmp := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return tmp
}