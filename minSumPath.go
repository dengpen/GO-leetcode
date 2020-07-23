package main

import "fmt"

func minPathSum(grid [][]int) int {

	// dp
	// dp = [][]int{}

	for i, v1 := range grid {
		for j, _ := range v1 {
			if j == 0 && i == 0 {
				continue
			}
			if i > 0 && j > 0 {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] = grid[i][j-1] + grid[i][j]
					continue
				} else {
					grid[i][j] = grid[i-1][j] + grid[i][j]
					continue
				}
			}
			if j == 0 && i > 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	slice := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	ret := minPathSum(slice)

	fmt.Printf("int: %v\n", ret)
}
