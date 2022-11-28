package link

import (
	"fmt"
	"testing"
)

var (
	linkA = &Node{
		Value: 5,
		Next: &Node{
			Value: 4,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 2,
					Next: &Node{
						Value: 1,
					},
				},
			},
		},
	}

	aaaaaa = &Node{
		Value: 5,
		Next: &Node{
			Value: 4,
			Next: &Node{
				Value: 3,
			},
		},
	}
)

func init() {
	linkA.Print("linkA")
	fmt.Println("init over")
}

func TestReverseList(t *testing.T) {
	n := ReverseList(linkA)
	n.Print("")
}

func TestSwapPairs(t *testing.T) {
	SwapPairs(linkA).Print("")
}

func TestHasCycle(t *testing.T) {
	// 构成环
	aaaaaa.Next.Next.Next = aaaaaa
	fmt.Println(HasCycle(aaaaaa))
}
