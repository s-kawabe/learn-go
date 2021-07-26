package main

import "fmt"

func lenAndCap() {
	s := []int{1, 2, 3, 4, 5}
	printSlice(s) // => len=5 cap=5 [1 2 3 4 5]

	s = s[:0]
	printSlice(s) // => len=0 cap=5 []

	s = s[:4]
	printSlice(s) // => len=4 cap=5 [1 2 3 4]

	s = s[2:]
	printSlice(s) // => len=2 cap=5 [3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
スライスのzero valuesはnil。
var s []int // => sはnil

nilスライスは0の長さと容量を持っており、何の元となる配列も持っていません。
*/
