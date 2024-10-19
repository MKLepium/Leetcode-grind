package main

type ListNode struct {
    Val int
    Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2;
	}
	if l2 == nil {
		return l1;
	}
	dummy := &ListNode{}
	current := dummy
	var overhang int
	var currentVal int

	for l1 != nil || l2 != nil || overhang != 0 {
		if l1 != nil {
			overhang += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			overhang += l2.Val
			l2 = l2.Next
		}
		currentVal = overhang % 10
		overhang /= 10
		current.Next = &ListNode{Val: currentVal}
		current = current.Next
	}
	return dummy.Next
}

func getNodes(a, b, c int) (int, int) {
	current := a + b + c
	if current > 9 {
		return current % 10, 1
	}
	return current, 0
}

func getNextNode(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}