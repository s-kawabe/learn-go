package main

import (
	"fmt"
	"strings"
)

func main2() {
	a := make([]int, 5) // nilではなく0で埋められたスライスを生成
	printSlice2("a", a)

	//        型    長さ 最大要素数
	b := make([]int, 0, 5)
	printSlice2("b", b) // len=5 cap=5 [0 0 0 0 0]

	c := b[:2]          // capは変わらない
	printSlice2("c", c) // len=2 cap=5 [0 0]

	d := c[2:5]         // cap変わる
	printSlice2("d", d) // len=3 cap=3 [0 0 0]
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// sliceの中にsliceを入れることができる
func twoSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
