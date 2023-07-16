- 递归
  - 递归函数有两个条件：基线条件、递归条件
- 快速排序
  - D&C将问题分解，使用D&C处理列表时，基线条件很可能是空数组或只包含一个元素的数组
  - 实现快速排序时，随机选择用作基准值的元素，快排平均运行时间为O(nlogn)
  - 大O表示法中的常量，是快排比合并排序快的原因所在
  - 比较简单查找和二分查找时，常量几乎无关紧要，因为列表很长时，O(logn)的速度比O(n)快得多
- 散列表
  - 适用于模拟映射关系，防止重复，缓存/记住数据，以免服务器再次生成它们
- 广度搜索优先
  - 常用于寻找最短路径，对于类似问题，可先用图建立模型，再使用广度搜索优先来解决问题
  - 队列先进先出(FIFO)，栈后进先出(LIFO)，需要**按加入顺序检查搜索列表的人，否则找到的就不是最短路径，所以搜索列表必须是队列**
  - 对于检查过的人，务必不要再检查，否则可能无限循环
- 狄克斯特拉算法
  - 广度优先算法用于在非加权图中寻找最短路径
  - 狄克斯特拉算法用于在加权图中寻找最短路径
  - 仅当权重为正时，狄克斯特拉算法才有用
  - 如果图中有负权重，应使用贝尔曼-福德算法
- 贝尔曼-福德算法(Bellman-Ford)
  - 与 Dijkstra 同为解决单源最短路径的算法，不同的是，Dijkstra采用贪婪算法范式进行设计，Bellman-Ford 采用动态规划设计，且能适应一般情况，即存在负权边的情况
  - 一个实现的很好的 Dijkstra 算法比 Bellman-Ford 算法的运行时间要低
- 贪婪算法
  - 贪婪算法用于寻找局部最优解
  - 对于NP完全问题，尚无快捷解决方案，最佳做法是使用近似算法
  - 贪婪算法易于实现、运行速度快，是不错的[近似算法](https://zh.m.wikipedia.org/zh-hans/%E8%BF%91%E4%BC%BC%E7%AE%97%E6%B3%95)
- 动态规划
  - 需要在给定约束条件下优化某种指标时，可使用动态规划
  - 问题可分解为离散子问题，即不依赖其它子问题时，可使用动态规划
  - 动态规划必然涉及网格，每个单元格都是一个子问题，因此需要考虑如何将问题分解为子问题
  - 动态规划没有通用公式
- K最邻近算法
  - k nearest neighbors(KNN)，用于分类和回归，需要考虑最近的邻居
  - 分类就是编组，回归就是预测结果（如数字）
  - 特征抽取是指将物品转换为一系列可比较的数字
  - KNN算法的关键在于提取合适的特征

### 线性表

线性表主要包含两种存储结构 —— 链表、数组。

**数组**

所有元素**连续**存储与一段内存中，且每个元素占用的**内存大小相同**。因此数组具备通过下表快速访问数据的能力。

但因此带来的缺点是，增删元素的成本很高，时间复杂度均为O(n)。增加数组容量需要申请新内存，然后复制所有元素，可能还要删除原先的内存。

删除元素时需要移动被删除元素之后的所有元素，以确保连续性。增加元素时需要移动指定位置及之后的所有元素，然后将新增元素插入指定位置，若容量不足还需要先扩容。

- 优点：根据偏移量实现快速读写
- 缺点：扩容、增删元素慢

**链表**

由若干结点组成，每个结点包含数据域和指针域。

这种存储方式，使其可以高效的在指定位置插入与删除，时间复杂度均为O(1)。

无法高效获取长度，无法根据偏移快速访问元素，是链表的两个主要缺点。因此，关于链表的常见问题大多围绕以上两点，例如**获取倒数第k个元素、获取中间位置元素、判断链表是否存在环、判断环的长度**等与长度、位置有关的话题。这些问题都可以通过灵活运用**双指针**的方式来解决。

```go
/*
倒数第k个元素的问题
设有两个指针 p 和 q，初始均指向头结点，
先让 p 沿着 next 移动 k 次，此时 p 指向第 k+1 个结点，q 指向头结点，两者相距 k，
然后同时移动 p 和 q，直到 p 指向空，此时 q 所指结点即为倒数第 k 个结点。
*/
func lastKth(head *ListNode, k int) *ListNode {
	p, q := head, head
	for k > 0 {
		p = p.next
		k--
	}
	for p != nil {
		p = p.next
		q = q.next
	}
	return q
}
```

```go
/*
获取中间元素的问题
设有两个指针 fast 和 slow，初始均指向头结点，
每次移动时，fast 向后走两步，slow 向后走一步，直到 fast 无法向后走两步为止，
这使得每轮移动后，fast 和 slow 的间距会增加1，
设链表有 n 个元素，那么最多移动 n/2 轮，
n 为奇数时，slow 恰好指向中间结点，n 为偶数时，slow 恰好指向中间两结点中考前的那个
*/
func middleNode(head *ListNode) *ListNode {
    p, q := head, head
    for q != nil && q.next != nil {
        p = p.next
        q = q.next.next
    }
    return p
}
```

```go
/*
是否存在环的问题
如果尾结点的 next 指针指向其他任一结点，则链表存在环。
当链表有环时，快慢指针会陷入环中无限次移动，变为追及问题，只要一直移动下去，快指针总会追上慢指针
*/
func hasRing(head *ListNode) bool {
    fast, slow := head, head
    for fast != nil {
        fast = fast.next
        if fast != nil {
            fast = fast.next
        }
        if fast == slow {
            return true
        }
        slow = slow.next
    }
    return false
}
```

```go
/*
环的入口点问题
*/
func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil {
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            fast = head
            for fast != slow {
                fast = fast.Next
                slow = slow.Next
            }
            return fast
        }
    }
    return nil
}
```



```go
/*
如果存在环，环的长度是多少
快慢指针相遇后继续移动，直到第二次相遇，两次相遇间的移动次数即为环的长度
*/
func lengthOfRing(head *ListNode) int {
    fast, slow := head, head
    meet, count := 0, 0
    for meet < 1{
        fast = fast.next.next
        if fast == slow {
            meet++
            continue
        }
        slow = slow.next
    }
    for meet < 2{
        fast = fast.next.next
        count++
        if fast == slow {
            meet++
            continue
        }
        slow = slow.next
    }
    return count
}
```



### 树的遍历

**深度优先遍历**

- 前序遍历
依序以根节点、左节点、右节点为顺序访问

```
    void pre_order_traversal(TreeNode *root) {
    // Do Something with root
    if (root->lchild != NULL) //若其中一側的子樹非空則會讀取其子樹
        pre_order_traversal(root->lchild);
    if (root->rchild != NULL) //另一側的子樹也做相同事
        pre_order_traversal(root->rchild);
}
```

- 中序遍历

依序以左节点、根节点、右节点为顺序访问

```
    void in_order_traversal(TreeNode *root) {
    if (root->lchild != NULL) //若其中一側的子樹非空則會讀取其子樹
        in_order_traversal(root->lchild);
    // Do Something with root
    if (root->rchild != NULL) //另一側的子樹也做相同事
        in_order_traversal(root->rchild);
}
```

- 后序遍历

依序以左节点、右节点、根节点为顺序访问

```
    void post_order_traversal(TreeNode *root) {
    if (root->lchild != NULL) //若其中一側的子樹非空則會讀取其子樹
        post_order_traversal(root->lchild);
    if (root->rchild != NULL) //另一側的子樹也做相同事
        post_order_traversal(root->rchild);
    // Do Something with root
}
```

**广度优先遍历**
优先访问离根节点最近的节点。二叉树的广度优先遍历又称为*按层次遍历*。算法借助队列实现。

```
  void level(TreeNode *node)
{
  Queue *queue = initQueue();
  enQueue(queue, node);

  while (!isQueueEmpty(queue))
  {
    TreeNode *curr = deQueue(queue);

    // Do Something with curr

    if (curr->lchild != NULL)
      enQueue(queue, curr->lchild);
    if (curr->rchild != NULL)
      enQueue(queue, curr->rchild);
  }
}
```

## Reference
[《图解算法》](https://book.douban.com/subject/26979890/)
[grokking-algorithm](https://github.com/egonSchiele/grokking_algorithms)
[《算法第四版》](https://book.douban.com/subject/19952400/)
[Algorithms, 4th Edition](https://algs4.cs.princeton.edu/home/)

