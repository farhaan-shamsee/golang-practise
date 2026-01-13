/*
Problem: Number of Islands
Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water),
return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are surrounded by water.

Real-world use case for Platform Engineers:
- Identifying isolated network segments in VPC architecture
- Finding disconnected clusters in distributed systems
- Detecting separate availability zones or regions
- Analyzing isolated service groups in microservices mesh
- Identifying connected component failures in infrastructure

Approaches:

1. DFS (Depth-First Search):
  - For each unvisited land cell, start DFS
  - Mark all connected lands as visited
  - Each DFS call represents one island
  - Time: O(m × n), Space: O(m × n) for recursion

2. BFS (Breadth-First Search):
  - Similar to DFS but uses queue
  - Time: O(m × n), Space: O(min(m, n)) for queue

3. Union-Find (Disjoint Set):
  - Union adjacent land cells
  - Count number of disjoint sets
  - Time: O(m × n × α(m×n)), Space: O(m × n)

Algorithm (DFS Approach):
 1. Initialize count = 0
 2. Iterate through each cell in grid:
    a. If cell is '1' (unvisited land):
    - Increment count (found new island)
    - Call DFS to mark entire island as visited
 3. DFS marks current cell and recursively visits all 4 neighbors
 4. Return count

Time Complexity: O(m × n) - visit each cell once
Space Complexity: O(m × n) - worst case recursion depth (all land)

Example: grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"],
	["0","0","0","1","1"]

]

Visualization:

	L L . . .    Island 1 (top-left)
	L L . . .
	. . L . .    Island 2 (middle)
	. . . L L    Island 3 (bottom-right)

Step-by-step:

 1. Start at (0,0) = '1' → count=1, DFS marks (0,0), (0,1), (1,0), (1,1)
    Grid after DFS:
    V V . . .  (V = visited)
    V V . . .
    . . 1 . .
    . . . 1 1

 2. Continue scanning, find (2,2) = '1' → count=2, DFS marks (2,2)
    Grid after DFS:
    V V . . .
    V V . . .
    . . V . .
    . . . 1 1

 3. Continue scanning, find (3,3) = '1' → count=3, DFS marks (3,3), (3,4)
    Grid after DFS:
    V V . . .
    V V . . .
    . . V . .
    . . . V V

4. No more '1's found → return count = 3

Result: 3 islands
*/
package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Helper function to perform DFS
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// Boundary check and water check
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}

		// Mark as visited by changing '1' to '0'
		grid[i][j] = '0'

		// Explore all 4 directions (up, down, left, right)
		dfs(i-1, j) // up
		dfs(i+1, j) // down
		dfs(i, j-1) // left
		dfs(i, j+1) // right
	}

	// Scan entire grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++   // Found new island
				dfs(i, j) // Mark entire island
			}
		}
	}

	return count
}

func mainNumOfIslands() {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Grid 1:")
	printGrid(grid1)
	fmt.Printf("Number of islands: %d\n\n", numIslands(grid1)) // 1

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Grid 2:")
	printGrid(grid2)
	fmt.Printf("Number of islands: %d\n\n", numIslands(grid2)) // 3

	// Real-world example: Network segments
	fmt.Println("Network segments (1=connected, 0=isolated):")
	network := [][]byte{
		{'1', '0', '1', '0'},
		{'0', '0', '1', '0'},
		{'1', '1', '0', '0'},
		{'0', '0', '0', '1'},
	}
	printGrid(network)
	fmt.Printf("Number of isolated network segments: %d\n", numIslands(network))
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
