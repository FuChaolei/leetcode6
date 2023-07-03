//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, preodd, preeven := head, &ListNode{Val: 0, Next: nil}, &ListNode{Val: 0, Next: nil}
	odd := preodd
	even := preeven
	for cur != nil && cur.Next != nil {
		odd.Next = cur
		even.Next = cur.Next
		cur = cur.Next.Next
		odd = odd.Next
		even = even.Next
		odd.Next = nil
		even.Next = nil
	}
	if cur != nil {
		odd.Next = cur
	}

	cur = preeven.Next
	for cur.Next != nil {
		tail := cur.Next
		cur.Next = tail.Next
		tail.Next = preeven.Next
		preeven.Next = tail
	}
	odd = preodd.Next
	even = preeven.Next
	dummy := ListNode{Val: 0}
	pre := &dummy

	for odd != nil && even != nil {
		if odd.Val < even.Val {
			pre.Next = odd
			odd = odd.Next
		} else {
			pre.Next = even
			even = even.Next
		}
		pre = pre.Next
	}
	if even != nil {
		pre.Next = even
	}
	if odd != nil {
		pre.Next = odd
	}
	return dummy.Next
}
func main() {
	dummy := &ListNode{}
	pre := dummy
	pre.Next = &ListNode{Val: 1}
	pre = pre.Next
	pre.Next = &ListNode{Val: 3}
	pre = pre.Next
	pre.Next = &ListNode{Val: 2}
	pre = pre.Next
	pre.Next = &ListNode{Val: 2}
	pre = pre.Next
	pre.Next = &ListNode{Val: 3}
	pre = pre.Next
	pre.Next = &ListNode{Val: 1}
	pre = pre.Next
	res := sortLinkedList(dummy.Next)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
