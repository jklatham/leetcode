// 2181 - Merge Nodes in Between Zeros
package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	cur := head
	dummy := new(ListNode)
	tail := dummy
	for cur.Next != nil {
		node := new(ListNode)

		// 0  1 2 3  0  4 5 6  0
		for cur.Next.Val != 0 {
			node.Val += cur.Next.Val
			cur = cur.Next
		}
		tail.Next = node
		tail = tail.Next
		cur = cur.Next
	}

	return dummy.Next
}
