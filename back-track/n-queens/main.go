package main

import (
	"fmt"
	"strings"
)

func main() {
	SolveQueen(8)

	for _, r := range Res{
		for _, s := range r {
			fmt.Println(s)
		}
		fmt.Println("---------")
	}
}

var Res [][]string

// 输入棋盘边长 N 返回所有合法的放置
func SolveQueen(n int) [][]string {
	// . 表示空
	// Q 表示皇后
	// 初始化空棋盘 board[row][col]
	var board []string
	row := strings.Repeat("X", n)
	for i := 0; i < n; i++ {
		board = append(board, row)
	}
	backtrack(&board, 0)
	return Res
}

func backtrack(board *[]string, row int) {
	// 结束条件，遍历完所有的 row
	if len(*board) == row {
		res := append([]string{}, *board...)
		Res = append(Res, res)
		return
	}

	// 一行的长度
	n := len((*board)[row])
	// 遍历这一行，一个格子一个格子尝试放置
	for col := 0; col < n; col++ {
		// 判断这个位置能不能放置，排除不合法选项
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		temp := []byte((*board)[row])
		temp[col] = 'Q'
		(*board)[row] = string(temp)
		// 进入下一层决策
		backtrack(board, row + 1)
		// 撤销选择
		temp = []byte((*board)[row])
		temp[col] = 'X'
		(*board)[row] = string(temp)
	}
}

func isValid(board *[]string, row int, col int) bool {
	n := len(*board)
	// 检查这一列从上到下列是否有 Queen 存在
	for i := 0; i < n; i++ {
		if []byte((*board)[i])[col] == 'Q' {
			return false
		}
	}
	// 因为是从上往下放置，只要检查右上方和左上方是否冲突即可

	// 检测右上方是否有皇后冲突
	for i, j := row - 1, col + 1 ; i >= 0 && j < n; i, j = i - 1, j + 1 {
		if []byte((*board)[i])[j] == 'Q' {
			return false
		}
	}

	// 检测左上方
	for i, j := row - 1, col - 1 ; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
		if []byte((*board)[i])[j] == 'Q' {
			return false
		}
	}
	return true
}
