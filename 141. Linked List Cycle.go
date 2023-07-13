```
思路1：
快慢指针，快指针一次走两步，慢指针一次走一步，如果存在环，则两指针一定会相遇
```

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    fast, slow := head, head
    for fast != nil {
        fast = fast.Next
        if fast != nil{
            fast = fast.Next
        }
        if fast == slow {
            return true
        }
        slow = slow.Next
    }
    return false
}

```
思路2：
哈希表，遍历所有节点，每次我们到达一个节点，如果该节点已经存在于哈希表中，则说明该链表是环形链表，否则就将该节点加入哈希表中，重复这一过程，直到遍历完整个链表即可。

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    m := map[*ListNode]struct{}{}
    for head != nil {
        if _, ok := m[head]; ok {
            return true
        }
        m[head] = struct{}{}
        head = head.Next
    }
    return false
}

```
思路3：
另辟蹊径，遍历每个节点，给每个访问过的节点赋特殊值，如果访问到与特殊值相等的节点，则有环
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    for head != nil {
        if head.Val == 100000 {
            return true
        }
        head.Val = 100000
        head = head.Next
    }
    return false
}