package stack

import "strings"

type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	tmp := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return tmp
}
func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

// leetcode 20
// 括号是否合法    O(n)
func IsVailed1(s string) bool {
	stack := []rune{}
	kuohaoM := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, item := range s {
		if v, ok := kuohaoM[rune(item)]; !ok {
			stack = append(stack, item)
		} else if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if tmp != v {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// 括号是否合法  0(n^2/2)
func IsVailed2(s string) bool {
	lenth := 0
	for lenth != len(s) {
		lenth = len(s)
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "()", "")
	}
	return len(s) == 0
}

// 只用栈实现一个队列，先进先出
type Stack2Queue struct {
	in, out Stack
}

func (s *Stack2Queue) Push(v int) {
	for _, v := range s.out.values {
		s.in.Push(v)
	}
	s.in.Push(v)
}

func (s *Stack2Queue) Pop() int {
	if !s.out.IsEmpty() {
		return s.out.Pop()
	}
	for !s.in.IsEmpty() {
		i := s.in.Pop()
		s.out.Push(i)
	}
	return s.out.Pop()
}
