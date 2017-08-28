# 2. Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

###### Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
###### Output: 7 -> 0 -> 8

解题思路：模拟横竖加减分法。[代码源地址](http://blog.csdn.net/neosmith/article/details/53302549)

> Tips：千万不要直接累加然后取余除10，这样容易溢出，而且效率也不是很高。既然提供了链表，就从链表的特性出发，简化计算。

```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promotion := 0

	var head *ListNode
	var rear *ListNode

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += promotion
		promotion = 0

		if sum >= 10 {
			promotion = 1
			sum = sum % 10
		}

		node := &ListNode{
			sum,
			nil,
		}

		if head == nil {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = node
		}

	}

	if promotion > 0 {
		rear.Next = &ListNode{
			promotion,
			nil,
		}
	}

	return head
}
```

