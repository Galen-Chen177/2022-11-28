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
			Left: &Tree{
				Value: 5,
			},
		},
	}
)

func TestInOrder(t *testing.T) {
	fmt.Println(InOrder(treeA))
}

func TestIsVailedBST(t *testing.T) {
	fmt.Println(IsVailedBST1(treeA))
	fmt.Println(IsVailedBST2(treeA))

}
