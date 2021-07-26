package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // => Hello World
	fmt.Println(a)          // [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // => [2 3 5 7 11 13]
}

/*
配列の長さは型の一部分である。
Goにおける"配列"はサイズを変えることはできない。 => 固定長
*/
