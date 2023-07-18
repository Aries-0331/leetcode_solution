/**
思路：快慢指针
1. 找到前半部分链表的尾结点
2. 反转后半部分链表
3. 判断是否回文
4. 恢复被反转的链表
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    half := findHalf(head)
    rev := reverseList(half.Next)
    p := head
    for q := rev; q != nil; q = q.Next {
        if p.Val != q.Val {
            return false
        }
        p = p.Next
    }
    half.Next = reverseList(rev)
    return true
}

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

func findHalf(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}