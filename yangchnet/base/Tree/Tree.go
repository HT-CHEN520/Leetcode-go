package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}
	rootIndex := indexOf(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:1+rootIndex], inorder[:rootIndex]),
		Right: BuildTree(preorder[1+rootIndex:], inorder[rootIndex+1:]),
	}
	return root
}

func (root *TreeNode) InOrder(Visit func(val int)) {
	if root == nil {
		return
	}
	root.Left.InOrder(Visit)
	Visit(root.Val)
	root.Right.InOrder(Visit)
}

func (root *TreeNode) PreOrder(Visit func(val int)) {
	if root == nil {
		return
	}
	Visit(root.Val)
	root.Left.PreOrder(Visit)
	root.Right.PreOrder(Visit)
}

func (root *TreeNode) PostOrder(Visit func(val int)) {
	if root == nil {
		return
	}
	root.Left.PostOrder(Visit)
	root.Right.PostOrder(Visit)
	Visit(root.Val)
}

func indexOf(order []int, val int) int {
	for i, v := range order {
		if val == v {
			return i
		}
	}
	return -1
}
