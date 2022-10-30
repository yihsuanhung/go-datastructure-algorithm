package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {

		size := len(queue)

		level := []int{}

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		result = append(result, level)

	}

	return result
}

func main() {}
