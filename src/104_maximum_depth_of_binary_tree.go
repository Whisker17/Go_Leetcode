package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	}
	//return int(math.Max(float64(maxDepth(root.Left)),float64(maxDepth(root.Right)))+1)
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
