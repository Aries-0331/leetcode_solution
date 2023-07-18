/*
思路1：
如果两个链表相交，那么相交点之后的长度是相同的。
所以，让两个链表从同距离末尾相同距离的位置开始遍历，这个位置只能是较短链表的头结点位置，
因此需要消除两个链表的长度差：
1. 指针p指向A链表，指针q指向B链表，依次往后遍历
2. 如果p到了末尾，则将p置为headB，继续遍历
3. 如果q到了末尾，则将q置为headA，继续遍历
4. 比较长的链表指针指向较短链表head时，长度差就消除了
5. 如此，只需要将最短链表遍历两次即可找到位置
空间复杂度O(1)，时间复杂度O(n)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    p, q := headA, headB
    for p != q {
        if p == nil {
            p = headB
        }else{
            p = p.Next
        }
        if q == nil {
            q = headA
        }else{
            q = q.Next
        }
    }
    return p
}

/*
思路2：
哈希集合
判断两链表是否相交，可使用哈希集合存储链表节点。
首先遍历链表headA，将每个节点存入哈希集合，
然后遍历链表headB，对于遍历的每个节点，判断该节点是否在哈希集合中，
若不在，继续遍历下一个，若在，则后面的节点都在哈希集合中，该节点即为相交节点。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := map[*ListNode]bool{}
    for p := headA; p != nil; p = p.Next {
        m[p] = true
    }
    for p := headB; p != nil; p = p.Next {
        if m[p] {
            return p
        }
    }
    return nil
}