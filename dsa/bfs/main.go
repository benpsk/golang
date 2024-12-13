package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.Value, " ")

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func main() {
	// Example Tree:
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   6
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}
	root.Right.Right = &TreeNode{Value: 6}

	fmt.Println("Breadth first search traversal:")
	bfs(root)
}
