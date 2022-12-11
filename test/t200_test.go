package test

import "testing"

func numIslands(grid [][]byte) int {
	m, n, res := len(grid), len(grid[0]), 0
	for i := 0; i < m*n; i++ {
		row, col := i/n, i%n
		if grid[row][col] == '0' {
			continue
		}
		//bfs(grid, i, m, n)
		dfs(grid, i, m, n)
		res++
	}
	return res
}

/**
按照右、下、左、上的顺序
*/
func bfs(grid [][]byte, i, m, n int) {
	queue := []int{i}
	for len(queue) > 0 {
		head := queue[0]
		row, col := head/n, head%n
		grid[row][col] = '0'
		// 向右
		if col+1 < n && grid[row][col+1] == '1' {
			grid[row][col+1] = '0'
			j := row*n + col + 1
			queue = append(queue, j)
		}
		// 向下
		if row+1 < m && grid[row+1][col] == '1' {
			grid[row+1][col] = '0'
			j := (row+1)*n + col
			queue = append(queue, j)
		}
		// 向左
		if col-1 >= 0 && grid[row][col-1] == '1' {
			grid[row][col-1] = '0'
			j := row*n + col - 1
			queue = append(queue, j)
		}
		// 向上
		if row-1 >= 0 && grid[row-1][col] == '1' {
			grid[row-1][col] = '0'
			j := (row-1)*n + col
			queue = append(queue, j)
		}
		queue = queue[1:]
	}
}

/**
按照右、下、左、上的顺序
*/
func dfs(grid [][]byte, i, m, n int) {
	row, col := i/n, i%n
	grid[row][col] = '0'
	// 向右
	if col+1 < n && grid[row][col+1] == '1' {
		grid[row][col+1] = '0'
		j := row*n + col + 1
		dfs(grid, j, m, n)
	}
	// 向下
	if row+1 < m && grid[row+1][col] == '1' {
		grid[row+1][col] = '0'
		j := (row+1)*n + col
		dfs(grid, j, m, n)
	}
	// 向左
	if col-1 >= 0 && grid[row][col-1] == '1' {
		grid[row][col-1] = '0'
		j := row*n + col - 1
		dfs(grid, j, m, n)
	}
	// 向上
	if row-1 >= 0 && grid[row-1][col] == '1' {
		grid[row-1][col] = '0'
		j := (row-1)*n + col
		dfs(grid, j, m, n)
	}
}

func TestDo200(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'},
		{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'},
		{'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'},
		{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '0'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'},
	}
	res := numIslands(grid)
	println(res)
}
