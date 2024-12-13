package main

import "fmt"

// Define the TreeNode structure
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Inorder Traversal
func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		fmt.Print(root.Value, " ")
		inorder(root.Right)
	}
}

// Preorder Traversal
func preorder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Value, " ")
		preorder(root.Left)
		preorder(root.Right)
	}
}

// Postorder Traversal
func postorder(root *TreeNode) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		fmt.Print(root.Value, " ")
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

	fmt.Println("Inorder Traversal:")
	inorder(root) // Output: 4 2 5 1 3 6
	fmt.Println()

	fmt.Println("Preorder Traversal:")
	preorder(root) // Output: 1 2 4 5 3 6
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	postorder(root) // Output: 4 5 2 6 3 1
	fmt.Println()
}
