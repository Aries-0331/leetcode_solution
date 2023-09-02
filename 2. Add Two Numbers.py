# 递归，边界为 l1 和 l2 都为空节点，如果进位了，则创建一个新节点
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode], carry=0) -> Optional[ListNode]:
        if l1 is None and l2 is None:
            return ListNode(carry) if carry else None
        if l1 is None:
            l1, l2 = l2, l1
        carry += l1.val + (l2.val if l2 else 0)
        l1.val = carry % 10
        l1.next = self.addTwoNumbers(l1.next, l2.next if l2 else None, carry // 10)
        return l1
    
# 迭代，迭代条件为 两个链表存在非空节点或进位值不为0
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cur = tmp = ListNode()
        carry = 0
        while l1 or l2 or carry:
            carry += (l1.val if l1 else 0) + (l2.val if l2 else 0)
            cur.next = ListNode(carry % 10)
            carry //= 10
            cur = cur.next
            if l1: l1 = l1.next
            if l2: l2 = l2.next
        return tmp.next