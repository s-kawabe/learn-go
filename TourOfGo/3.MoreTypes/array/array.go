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

func sub() {
	// 要素数の省略
	a := [...]int{1, 2, 3} // a => [1, 2, 3]

	// 要素数を...で省略して書くことができる。
	// この場合の要素数の初期値は与えたものの数が初期値となる。
	// 型として [...]int と認識されるわけではない。
}

/*
配列の長さは型の一部分である。
Goにおける"配列"はサイズを変えることはできない。 => 固定長
*/
