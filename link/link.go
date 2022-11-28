package link

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Print(info string) {
	if info != "" {
		fmt.Printf("%s\t", info)
	}
	cur := n
	for cur != nil {
		fmt.Printf("%d\t", cur.Value)
		cur = cur.Next
	}
	fmt.Println()
}

// 反转链表   O(n)
func ReverseList(head *Node) *Node {
	cur := head
	var prev *Node
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}

// 排序,升序或降序
func SortLink(head *Node, up bool) *Node {
	panic("")
}

// 链表两两交换
// 1->2->3->4    ==> 2->1->4->3
// 1->2->3->4->5 ==> 2->1->4->3->5
func SwapPairs(head *Node) *Node {
	res := &Node{}
	pre := res
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := pre.Next.Next
		a.Next = b.Next
		b.Next = a
		pre.Next = b
		pre = a
	}
	return res.Next
}

// 是否有环
func HasCycle(head *Node) bool {
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
