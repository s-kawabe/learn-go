package main

import "fmt"

func split(sum int) (x, y int) {
	// 既にxとyは初期化されている
	x = sum * 4 / 9
	y = sum - x
	return // xとyを返す
}

func main() {
	fmt.Println(split(17))
}

/*
戻り値となる変数に名前をつけることができる。
これを指定すると、return　のみでも記載した変数を返すことができる。 → (naked return)
指定した変数は、関数内の最初に定義されたと同等になる。

→長い関数で使用すると可読性が下がる！！
*/
