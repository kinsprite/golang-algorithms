package search

//
//       << LeftRotate
//          RigthRotate >>
//       Y                X
//    X     r          a     Y
//  a   b                  b   r
//

// RbLeftRotate Left Rotate
func RbLeftRotate(T *Tree, x *TreeNode) {
	y := x.Right

	// move y.Left to x.Right
	x.Right = y.Left

	if y.Left != TreeNil {
		y.Left.Parent = x
	}

	// move y to x position
	y.Parent = x.Parent

	if x.Parent == TreeNil {
		T.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	// update x's new position
	y.Left = x
	x.Parent = y
}

// RbRightRotate RightRotate
func RbRightRotate(T *Tree, y *TreeNode) {
	x := y.Left

	// move x.Right to y.Left
	y.Left = x.Right

	// move x to y position
	if y.Parent == TreeNil {
		T.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	// update y's new postion
	x.Right = y
	y.Parent = x
}

// RbTreeInsert RbTreeInsert
func RbTreeInsert(T *Tree, z *TreeNode) {
	y := TreeNil
	x := T.Root

	// 找出待插入的父节点
	for x != TreeNil {
		y = x

		if z.Key < y.Key {
			x = y.Left
		} else {
			x = y.Right
		}
	}

	z.Parent = y

	// 在父节点中挂载z
	if y == TreeNil {
		T.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}

	z.Left = TreeNil
	z.Right = TreeNil
	z.Color = Red

	RbInsertFixup(T, z)
}

// RbInsertFixup RbInsertFixup
func RbInsertFixup(T *Tree, z *TreeNode) {
	for z.Parent.Color == Red {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right

			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else if z == z.Parent.Right {
				z = z.Parent
				RbLeftRotate(T, z.Parent.Parent)
			} else {
				z.Parent.Color = Black
				z.Parent.Parent.Color = Red
				RbRightRotate(T, z.Parent.Parent)
			}

		} else {
			y := z.Parent.Parent.Left

			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else if z == z.Parent.Left {
				z = z.Parent
				RbRightRotate(T, z.Parent.Parent)
			} else {
				z.Parent.Color = Black
				z.Parent.Parent.Color = Red
				RbLeftRotate(T, z.Parent.Parent)
			}
		}

		T.Root.Color = Black
	}
}

// RbTreeDelete RbTreeDelete
func RbTreeDelete(T *Tree, z *TreeNode) {
	y := z
	yOriginalColor := y.Color
	var x *TreeNode

	if z.Left == TreeNil {
		x = z.Right
		RbTreeTransplant(T, z, z.Right)
	} else if z.Right == TreeNil {
		x = z.Left
		RbTreeTransplant(T, z, z.Left)
	} else {
		y := TreeMinimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			x.Parent = y
		} else {
			// 替换y位置
			TreeTransplant(T, y, y.Right)
			// 挂上右子树
			y.Right = z.Right
			y.Right.Parent = y
		}

		// 替换z位置
		RbTreeTransplant(T, z, y)
		// 挂上左子树
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == Black {
		RbTreeDeleteFixup(T, x)
	}
}

// RbTreeTransplant RbTreeTransplant
func RbTreeTransplant(T *Tree, u, v *TreeNode) {
	if u.Parent == TreeNil {
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

// RbTreeDeleteFixup RbTreeDeleteFixup
func RbTreeDeleteFixup(T *Tree, x *TreeNode) {
	for x != T.Root && x.Color == Black {
		if x == x.Parent.Left {
			w := x.Parent.Right

			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				RbLeftRotate(T, x.Parent)
				w = x.Parent.Right
			}

			if w.Left.Color == Black && w.Right.Color == Black {
				w.Color = Red
				x = x.Parent
			} else if w.Right.Color == Black {
				w.Left.Color = Black
				w.Color = Red
				RbRightRotate(T, w)
				w = x.Parent.Right
			} else {
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				w.Right.Color = Black
				RbLeftRotate(T, x.Parent)
				x = T.Root
			}
		} else {
			w := x.Parent.Left

			if w.Color == Red {
				w.Color = Black
				x.Parent.Color = Red
				RbRightRotate(T, x.Parent)
				w = x.Parent.Right
			}

			if w.Left.Color == Black && w.Right.Color == Black {
				w.Color = Red
				x = x.Parent
			} else if w.Left.Color == Black {
				w.Right.Color = Black
				w.Color = Red
				RbLeftRotate(T, w)
				w = x.Parent.Left
			} else {
				w.Color = x.Parent.Color
				x.Parent.Color = Black
				w.Left.Color = Black
				RbRightRotate(T, x.Parent)
				x = T.Root
			}
		}
	}

	x.Color = Black
}
