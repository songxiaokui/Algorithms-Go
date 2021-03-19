package homework

/*
@Time    : 2021/3/19 14:31
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 200.numsOfLands.go
@Software: GoLand
*/

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

*/

func numIslands(grid [][]byte) int {
	// use dfs， if visit number is '1', count ++ then Smash up, down, left and
	//right '1'， until index out of range or visit '0'
	var result int
	for row := range grid {
		for col := range grid[row] {
			// if visit data is '1', result ++
			if grid[row][col] == '1' {
				result++
				// visit this point around point
				dfs(grid, row, col)
			}
		}
	}
	return result

}

func dfs(grid [][]byte, row, col int) {
	// terminal
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == '0' {
		return
	}

	// current process logic
	grid[row][col] = '0'
	// drill down

	// left
	dfs(grid, row, col-1)
	// right
	dfs(grid, row, col+1)
	// up
	dfs(grid, row-1, col)
	// down
	dfs(grid, row+1, col)
	// restore current status
}

func numIslands2(grid [][]byte) int {
	var result int
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '0' {
				continue
			}
			bfs(grid, row, col)
			result++
		}
	}
	return result
}

// bfs,use a queue to save un visit point, until not point
func bfs(grip [][]byte, row, col int) {
	var queue [][]int
	// add current point
	queue = append(queue, []int{row, col})
	// loop
	for len(queue) != 0 {
		// get first point location
		point := queue[0]
		// cut down
		queue = queue[1:]
		if point[0] >= 0 && point[0] < len(grip) && point[1] >= 0 &&
			point[1] < len(grip[point[0]]) && grip[point[0]][point[1]] == '1' {
			// update this point is '0'
			grip[point[0]][point[1]] = '0'
			// add this round point to queue
			queue = append(queue, []int{point[0], point[1] - 1},
				[]int{point[0], point[1] + 1}, []int{point[0] + 1, point[1]},
				[]int{point[0] - 1, point[1]})
		}
	}

}
