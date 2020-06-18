package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	g := [][]int{
		{1, 3, 1},
		{1, 5, 1},
	}
	fmt.Println(minPathSum(g))
}
func minPathSum(grid [][]int) int {
	maxX := len(grid)
	if maxX == 0 {
		return 0
	}
	maxY := len(grid[0])
	move := [][]int{{1, 0}, {0, 1}}
	min := 0
	getSum(0, move, grid, maxX, maxY, 0, 0, &min, []int{})
	return min
}

func getSum(sum int, m, grid [][]int, X, Y int, x, y int, min *int, path []int) {
	sum += grid[x][y]
	path = append(path, grid[x][y])
	if x == X-1 && y == Y-1 {
		if *min == 0 || sum < *min {
			*min = sum
		}
		fmt.Println(path)
		return
	}
	for _, v := range m {
		if x+v[0] < X && y+v[1] < Y {
			getSum(sum, m, grid, X, Y, x+v[0], y+v[1], min, path)
		}
	}
}
