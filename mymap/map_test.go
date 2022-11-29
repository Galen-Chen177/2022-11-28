package mymap

import (
	"fmt"
	"testing"
)

func TestTwoNums(t *testing.T) {
	nums := []int{0, 1, -1, 3, 5, 6}
	fmt.Println(TwoNums(nums, 8))
}

func TestThreeNums(t *testing.T) {
	nums := []int{-2, -1, 0, 1, 2, 3, 4, 5}
	fmt.Println(ThreeNums(nums, 0))
}
