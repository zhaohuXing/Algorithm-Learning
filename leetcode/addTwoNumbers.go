package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//addwoNumbers(l1, l2)
}

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
