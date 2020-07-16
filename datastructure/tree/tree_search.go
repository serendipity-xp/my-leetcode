package tree

/*
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。
	思路：
		一：中序遍历，检查结果列表是否已经有序
		二:分治法，判断左MAX < 根 < 右MIN
*/

// v1
func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	// check order
	for i := 0; i < len(result)-1; i++ {
		if result[i] > result[i+1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// v2
type BSTResultType struct {
	IsValid bool
	// 记录左右两边最大最小值，和根节点进行比较
	Max *TreeNode
	Min *TreeNode
}

func isValidBST2(root *TreeNode) bool {
	result := isValidBSTHelper(root)
	return result.IsValid
}

func isValidBSTHelper(root *TreeNode) BSTResultType {
	result := BSTResultType{}
	// check
	if root == nil {
		result.IsValid = true
		return result
	}
	// divide
	left := isValidBSTHelper(root.Left)
	right := isValidBSTHelper(root.Right)
	// conquer
	if !left.IsValid || !right.IsValid{
		result.IsValid = false
		return result
	}
	if left.Max != nil && left.Max.Val >= root.Val{
		result.IsValid = false
		return result
	}
	if right.Min != nil && right.Min.Val <= root.Val{
		result.IsValid = false
		return result
	}

	result.IsValid = true
	// 如果左（右）边还有更小的，就用更小的节点
	result.Min = root
	if left.Min != nil{
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil{
		result.Max = right.Max
	}
	return result
}

// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。返回插入后二叉搜索树的根节点。
// 思路：找到最后一个叶子节点满足插入条件即可。
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		root = &TreeNode{
			Val:   val,
		}
		return root
	}
	if root.Val > val{
		root.Left = insertIntoBST(root.Left, val)
	}else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
