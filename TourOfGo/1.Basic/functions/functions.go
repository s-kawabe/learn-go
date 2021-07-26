package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 20))
}

/*
関数は、0個以上の引数を取ることができます。
型は変数の後に記述する。
*/

func add2(x, y int) int {
	return x + y
}

/* 型の省略記法
型の2つ以上の引数が同じ型である場合には上記のような記述ができる
*/
