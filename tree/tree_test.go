package tree

import (
	"fmt"
	"testing"
)

var (
	treeA = &Tree{
		Value: 4,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 1,
			},
			Right: &Tree{
				Value: 3,
			},
		},
		Right: &Tree{
			Value: 6,
		},
	}
)

func TestInOrder(t *testing.T) {
	fmt.Println(InOrder(treeA))
}

// leetcode 98
func TestIsVailedBST(t *testing.T) {
	fmt.Println(IsVailedBST1(treeA))
	fmt.Println(IsVailedBST2(treeA))
}

// leetcode 102
func TestLeaveOrder(t *testing.T) {
	fmt.Println(LeaveOrderBFS(treeA))
	fmt.Println(LeaveOrderDFS(treeA))
}

// leetcode 104  111
func TestMaxDepthBFS(t *testing.T) {
	fmt.Println(MaxDepthBFS(treeA))
	fmt.Println(MaxDepthDFS(treeA))

}
