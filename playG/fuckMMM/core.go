package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := [][]byte{{'a', 'b'}}
	t := "ba"
	fmt.Println(exist(a, t))
}
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, index, x, y int) bool {
	fmt.Println(x, y)
	if x >= len(board) || x < 0 || y >= len(board[0]) || y < 0 {
		return false
	}
	if board[x][y] == ' ' {
		return false
	}
	if board[x][y] != word[index] {
		return false
	}

	if index == len(word)-1 {
		return true
	}
	tmp := board[x][y]
	board[x][y] = ' '
	res := dfs(board, word, index+1, x+1, y) ||
		dfs(board, word, index+1, x, y+1) ||
		dfs(board, word, index+1, x-1, y) ||
		dfs(board, word, index+1, x, y-1)
	board[x][y] = tmp
	return res
}
