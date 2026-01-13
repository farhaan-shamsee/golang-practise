/*
Problem: Maximum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node
down to the farthest leaf node.

Real-world use case for Platform Engineers:
- Navigating hierarchical resource ownership trees in cloud environments
- Analyzing nested configuration structures (YAML/JSON depth)
- Measuring depth of service dependency chains
- Evaluating organization hierarchy depth in access control
- Analyzing nested namespace structures in Kubernetes

Approaches:

1. Recursive DFS (Depth-First Search):
  - Base case: null node has depth 0
  - Recursive case: 1 + max(leftDepth, rightDepth)
  - Time: O(n), Space: O(h) where h = height (recursion stack)

2. Iterative BFS (Breadth-First Search):
  - Use queue for level-order traversal
  - Count levels as you traverse
  - Time: O(n), Space: O(w) where w = max width

3. Iterative DFS with Stack:
  - Use stack to simulate recursion
  - Track depth at each node
  - Time: O(n), Space: O(h)

Algorithm (Recursive DFS):
1. Base case: if root is null, return 0
2. Recursively find depth of left subtree
3. Recursively find depth of right subtree
4. Return 1 + max(leftDepth, rightDepth)
  - The "+1" accounts for current node

Time Complexity: O(n) - visit each node once
Space Complexity: O(h) - recursion stack depth = tree height
  - Best case (balanced): O(log n)
  - Worst case (skewed): O(n)

Example: Tree structure

	  3
	 / \
	9  20
	   / \
	  15  7

Visual representation with depths:

	  3      depth 1
	 / \
	9  20    depth 2
	   / \
	  15  7  depth 3

Execution trace:
maxDepth(3):

	├─ leftDepth = maxDepth(9):
	│   ├─ leftDepth = maxDepth(null) = 0
	│   └─ rightDepth = maxDepth(null) = 0
	│   └─ return 1 + max(0, 0) = 1
	│
	└─ rightDepth = maxDepth(20):
	    ├─ leftDepth = maxDepth(15):
	    │   ├─ leftDepth = maxDepth(null) = 0
	    │   └─ rightDepth = maxDepth(null) = 0
	    │   └─ return 1 + max(0, 0) = 1
	    │
	    └─ rightDepth = maxDepth(7):
	        ├─ leftDepth = maxDepth(null) = 0
	        └─ rightDepth = maxDepth(null) = 0
	        └─ return 1 + max(0, 0) = 1

	    └─ return 1 + max(1, 1) = 2

	└─ return 1 + max(1, 2) = 3

Result: Maximum depth = 3
*/
package main

import "fmt"

// TreeNode represents a node in binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive DFS approach
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to print tree (in-order traversal)
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(root.Val)

	newPrefix := prefix
	if isLeft {
		newPrefix += "│   "
	} else {
		newPrefix += "    "
	}

	if root.Left != nil || root.Right != nil {
		if root.Left != nil {
			printTree(root.Left, newPrefix, true)
		} else {
			fmt.Println(newPrefix + "├── null")
		}

		if root.Right != nil {
			printTree(root.Right, newPrefix, false)
		} else {
			fmt.Println(newPrefix + "└── null")
		}
	}
}

func mainMaxDepth() {
	// Example 1: Tree [3,9,20,null,null,15,7]
	//        3
	//       / \
	//      9  20
	//         / \
	//        15  7
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Tree 1:")
	printTree(root1, "", false)
	fmt.Printf("Maximum depth: %d\n\n", maxDepth(root1)) // 3

	// Example 2: Tree [1,null,2]
	//     1
	//      \
	//       2
	root2 := &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 2}

	fmt.Println("Tree 2:")
	printTree(root2, "", false)
	fmt.Printf("Maximum depth: %d\n\n", maxDepth(root2)) // 2

	// Example 3: Empty tree
	fmt.Println("Tree 3: (empty)")
	fmt.Printf("Maximum depth: %d\n\n", maxDepth(nil)) // 0

	// Real-world example: Service dependency tree
	fmt.Println("Service dependency tree:")
	fmt.Println("(API Gateway → Auth Service → Database)")
	serviceTree := &TreeNode{Val: 1}          // API Gateway
	serviceTree.Left = &TreeNode{Val: 2}      // Auth Service
	serviceTree.Left.Left = &TreeNode{Val: 3} // Database
	serviceTree.Right = &TreeNode{Val: 4}     // Cache Service

	printTree(serviceTree, "", false)
	fmt.Printf("Maximum dependency chain depth: %d\n", maxDepth(serviceTree))
}
