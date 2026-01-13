/*
Problem: Clone Graph
Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph.

Each node in the graph contains:
- val: an integer value
- neighbors: a list of references to other nodes

Real-world use case for Platform Engineers:
- Duplicating infrastructure environments (dev → staging → prod)
- Creating clones of service dependency maps in multi-tenant platforms
- Copying network topology configurations
- Replicating microservices architecture blueprints
- Backup and disaster recovery scenarios

Approaches:

1. Depth-First Search (DFS) with Hash Map:
  - Use recursion to traverse graph
  - Use hash map to track visited/cloned nodes
  - Time: O(N + E) where N = nodes, E = edges
  - Space: O(N) for recursion stack and hash map

2. Breadth-First Search (BFS) with Hash Map:
  - Use queue for level-order traversal
  - Use hash map to track visited/cloned nodes
  - Time: O(N + E)
  - Space: O(N) for queue and hash map

Algorithm (DFS Approach):
 1. Create hash map to store: original node → cloned node
 2. Start DFS from given node:
    a. If node is null, return null
    b. If node already cloned (exists in map), return cloned node
    c. Create new node with same value
    d. Add to map: original → clone
    e. Recursively clone all neighbors
    f. Return cloned node
 3. Hash map ensures each node is cloned exactly once

Time Complexity: O(N + E) - visit each node and edge once
Space Complexity: O(N) - hash map stores N nodes, recursion depth ≤ N

Example:
Original Graph:

	1 ---- 2
	|      |
	4 ---- 3

Adjacency List:

	1: [2, 4]
	2: [1, 3]
	3: [2, 4]
	4: [1, 3]

DFS Traversal (starting from node 1):

Step 1: Clone node 1

	visited = {1: Node(1)}
	Clone neighbors of 1: [2, 4]

Step 2: Clone node 2 (from 1's neighbors)

	visited = {1: Node(1), 2: Node(2)}
	Clone neighbors of 2: [1, 3]
	- Node 1 already cloned, use existing

Step 3: Clone node 3 (from 2's neighbors)

	visited = {1: Node(1), 2: Node(2), 3: Node(3)}
	Clone neighbors of 3: [2, 4]
	- Node 2 already cloned, use existing

Step 4: Clone node 4 (from 3's neighbors)

	visited = {1: Node(1), 2: Node(2), 3: Node(3), 4: Node(4)}
	Clone neighbors of 4: [1, 3]
	- Both already cloned, use existing

Result: Complete deep copy with all connections preserved

Key Insight:
Hash map prevents infinite loops and ensures each node is cloned exactly once.
All references in the cloned graph point to cloned nodes, not original nodes.
*/
package main

import "fmt"

// Node represents a graph node
type Node struct {
	Val       int
	Neighbors []*Node
}

// Clone graph using DFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Map: original node → cloned node
	visited := make(map[*Node]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[*Node]*Node) *Node {
	// If already cloned, return the clone
	if clone, exists := visited[node]; exists {
		return clone
	}

	// Create clone of current node
	clone := &Node{Val: node.Val}
	visited[node] = clone

	// Clone all neighbors recursively
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, visited))
	}

	return clone
}

// Helper function to print graph (BFS)
func printGraph(node *Node) {
	if node == nil {
		fmt.Println("Empty graph")
		return
	}

	visited := make(map[*Node]bool)
	queue := []*Node{node}
	visited[node] = true

	fmt.Println("Graph adjacency list:")
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("Node %d: neighbors = [", current.Val)
		for i, neighbor := range current.Neighbors {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(neighbor.Val)

			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		fmt.Println("]")
	}
}

func mainCloneGraph() {
	// Create graph: 1 -- 2
	//               |    |
	//               4 -- 3
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	fmt.Println("Original Graph:")
	printGraph(node1)

	// Clone the graph
	cloned := cloneGraph(node1)

	fmt.Println("\nCloned Graph:")
	printGraph(cloned)

	// Verify they're different objects
	fmt.Printf("\nOriginal node1 address: %p\n", node1)
	fmt.Printf("Cloned node1 address: %p\n", cloned)
	fmt.Printf("Are they different objects? %v\n", node1 != cloned)
}
