package tree

import (
	"math"
)

type Tree struct {
	Value       int
	Left, Right *Tree
}

// 中序遍历
func InOrder(root *Tree) []int {
	if root == nil {
		return nil
	}
	return append(append(InOrder(root.Left), root.Value), InOrder(root.Right)...)
}

// 判断是否为二叉搜索树
// 解法1：着眼于整个树，中序遍历，查看是否升序即可
// 解法2：着眼于最小的数，看左子树是否小于根节点。。。。，需要递归
// 解法3：着眼于最终结点，设置一些限制，比如最小应该大于多少，最大应该大于多少等，需要递归
func IsVailedBST1(root *Tree) bool {
	// 中序遍历--> 得到一个切片--->判断是否为升序
	i := InOrder(root)
	flag := i[0]
	for _, v := range i {
		if v < flag {
			return false
		}
		flag = v
	}
	return true
}

func IsVailedBST2(root *Tree) bool {

	//return helper1(root)
	//return helper2(root)
	return helper3(root, math.MinInt, math.MaxInt)
}

func helper1(root *Tree) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Left.Value < root.Value {
			if helper1(root.Left) {
				if root.Right != nil {
					if root.Right.Value > root.Value {
						return helper1(root.Right)
					} else {
						return false
					}
				}
			}
		} else {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Value > root.Value {
			return helper1(root.Right)
		} else {
			return false
		}
	}

	return true
}

// helper1 加强版
func helper2(root *Tree) bool {
	if root == nil {
		return true
	}
	lFalg := false
	rFalg := false

	if root.Left != nil {
		if root.Value > root.Left.Value {
			lFalg = helper2(root.Left)
		}
	} else {
		lFalg = true
	}
	if root.Right != nil {
		if root.Value < root.Right.Value {
			rFalg = helper2(root.Right)
		}
	} else {
		rFalg = true
	}
	return lFalg && rFalg
}

func helper3(root *Tree, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Value <= min || root.Value >= max {
		return false
	}
	return helper3(root.Left, min, root.Value) && helper3(root.Right, root.Value, max)
}

// leetcode 236
// 最近公共祖先节点
func LowestCommonAncestor(root, p, q *Tree) *Tree {
	if root == nil {
		return root
	}
	if root == q || root == p {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil {
		// 左边有p或者q
		if right != nil {
			// 右边有p或者q
			return root
		} else {
			return left
		}
	} else {
		return right
	}
}

// BST 是有顺序的，左边永远小于根节点
func LowestCommonAncestorBST(root, p, q *Tree) *Tree {
	if root.Value > p.Value && root.Value > q.Value {
		return LowestCommonAncestorBST(root.Left, p, q)
	}
	if root.Value < q.Value && root.Value < p.Value {
		return LowestCommonAncestorBST(root.Right, p, q)
	}
	return root
}

// BFS 模板
func LeaveOrderBFS(root *Tree) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}

	visited := map[*Tree]struct{}{}
	queue := []*Tree{}
	queue = append(queue, root)

	for len(queue) != 0 {
		tmp := []int{}
		flag := len(queue)
		for _, v := range queue {
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = struct{}{}
			tmp = append(tmp, v.Value)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[flag:]
		res = append(res, tmp)
	}
	return res
}

var (
	resLeaveOrderDFS = [][]int{}
	resMaxDepthDFS   = [2]int{}
)

func LeaveOrderDFS(root *Tree) [][]int {
	if root == nil {
		return [][]int{}
	}
	_LeaveOrderDFS(root, 0)

	return resLeaveOrderDFS
}

func _LeaveOrderDFS(root *Tree, leave int) {
	if root == nil {
		return
	}
	if len(resLeaveOrderDFS) < leave+1 {
		resLeaveOrderDFS = append(resLeaveOrderDFS, []int{})
	}
	resLeaveOrderDFS[leave] = append(resLeaveOrderDFS[leave], root.Value)

	_LeaveOrderDFS(root.Left, leave+1)
	_LeaveOrderDFS(root.Right, leave+1)
}

// leetcode 104
func MaxDepthBFS(root *Tree) [2]int {
	if root == nil {
		return [2]int{}
	}
	visited := map[*Tree]struct{}{}
	queue := []*Tree{}
	queue = append(queue, root)
	min, max := 0, 0
	for len(queue) != 0 {
		tmpLen := len(queue)
		max++
		for _, v := range queue {
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = struct{}{}
			if v.Left == nil && v.Right == nil && min == 0 {
				min = max
			}

			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[tmpLen:]
	}
	return [2]int{min, max}
}

// leetcode 111
func MaxDepthDFS(root *Tree) [2]int {
	if root == nil {
		return [2]int{}
	}
	_MaxDepthDFS(root, 1)
	return resMaxDepthDFS

}

func _MaxDepthDFS(root *Tree, leave int) {
	if root == nil {
		return
	}
	if leave > resMaxDepthDFS[1] {
		resMaxDepthDFS[1] = leave
	}
	if root.Left == nil && root.Right == nil && (resMaxDepthDFS[0] > leave || resMaxDepthDFS[0] == 0) {
		resMaxDepthDFS[0] = leave
	}
	if root.Left != nil {
		_MaxDepthDFS(root.Left, leave+1)
	}
	if root.Right != nil {
		_MaxDepthDFS(root.Right, leave+1)
	}
}
