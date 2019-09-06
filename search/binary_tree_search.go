package search

import "fmt"

// TreeNode Tree Node
type TreeNode struct {
	Key    int
	Data   interface{}
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

// Tree Tree
type Tree struct {
	Root *TreeNode
	Size int
}

// InOrderTreeWalk In-order tree walk
func InOrderTreeWalk(x *TreeNode) {
	if x == nil {
		return
	}

	InOrderTreeWalk(x.Left)
	fmt.Print(x.Data, ' ')
	InOrderTreeWalk(x.Right)
}

// PreOrderTreeWalk Pre-order tree walk
func PreOrderTreeWalk(x *TreeNode) {
	if x == nil {
		return
	}

	fmt.Print(x.Data, ' ')
	PreOrderTreeWalk(x.Left)
	PreOrderTreeWalk(x.Right)
}

// PostOrderTreeWalk Post-order tree walk
func PostOrderTreeWalk(x *TreeNode) {
	if x == nil {
		return
	}

	PostOrderTreeWalk(x.Left)
	PostOrderTreeWalk(x.Right)
	fmt.Print(x.Data, ' ')
}

// TreeSearch Tree search
func TreeSearch(x *TreeNode, k int) *TreeNode {
	if x == nil || k == x.Key {
		return x
	}

	if k < x.Key {
		return TreeSearch(x.Left, k)
	}

	return TreeSearch(x.Right, k)
}

// IterativeTreeSearch IterativeTreeSearch
func IterativeTreeSearch(x *TreeNode, k int) *TreeNode {
	for x != nil && k != x.Key {
		if k < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	return x
}

// TreeMinimum TreeMinimum
func TreeMinimum(x *TreeNode) *TreeNode {
	for x.Left != nil {
		x = x.Left
	}

	return x
}

// TreeMaximum TreeMaximum
func TreeMaximum(x *TreeNode) *TreeNode {
	for x.Right != nil {
		x = x.Right
	}

	return x
}

/*
TreeSuccessor TreeSuccessor 中序输出x的后续节点
    5
  3   7
 2 4  6 8
1
例如，x = 7 4 6 节点
*/
func TreeSuccessor(x *TreeNode) *TreeNode {
	for x.Right != nil {
		return TreeMinimum(x.Right)
	}

	y := x.Parent

	// 第一个x为y的左节点
	for y != nil && y.Right == x {
		x = y
		y = y.Parent
	}

	return y
}

/*
TreePredecessor TreePredecessor 中序输出x的前面节点
    5
  3   7
 2 4  6 8
1

例如，x = 2 4 6 节点
*/
func TreePredecessor(x *TreeNode) *TreeNode {
	for x.Left != nil {
		return TreeMaximum(x.Left)
	}

	y := x.Parent

	// 第一个x为y的右节点
	for y != nil && y.Left == x {
		x = y
		y = y.Parent
	}

	return y
}

// TreeInsert TreeInsert
func TreeInsert(T *Tree, z *TreeNode) {
	var y *TreeNode
	x := T.Root

	// 找出待插入的父节点
	for x != nil {
		y = x

		if z.Key < y.Key {
			x = y.Left
		} else {
			x = y.Right
		}
	}

	z.Parent = y

	// 在父节点中挂载z
	if y == nil {
		T.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

// TreeDelete TreeDelete
func TreeDelete(T *Tree, z *TreeNode) {
	if z.Left == nil {
		TreeTransplant(T, z, z.Right)
	} else if z.Right == nil {
		TreeTransplant(T, z, z.Left)
	} else {
		y := TreeMinimum(z.Right)

		if y.Parent != z {
			// 替换y位置
			TreeTransplant(T, y, y.Right)
			// 挂上右子树
			y.Right = z.Right
			y.Right.Parent = y
		}

		// 替换z位置
		TreeTransplant(T, z, y)
		// 挂上左子树
		y.Left = z.Left
		y.Left.Parent = y
	}
}

// TreeTransplant 将v替换u节点
func TreeTransplant(T *Tree, u, v *TreeNode) {
	if u.Parent == nil {
		T.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}
