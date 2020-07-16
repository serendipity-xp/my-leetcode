package tree

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelResult := make([]int, 0)
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			levelResult = append(levelResult, currentNode.Val)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		result = append(result, levelResult)
	}
	return result
}

// 自底向上 层序遍历
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	// 反转结果
	reverse(result)
	return result
}

func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 锯齿形层次遍历，Z字形遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggle := false
	for len(queue) > 0 {
		levelResult := make([]int, 0)
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			levelResult = append(levelResult, currentNode.Val)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
			if toggle {
				reverseSlice(levelResult)
			}
		}
		result = append(result, levelResult)
		toggle = !toggle
	}
	return result
}

func reverseSlice(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		t := nums[i]
		nums[i] = nums[len(nums)-1-i]
		nums[len(nums)-1-i] = t
	}
}
