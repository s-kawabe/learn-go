package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
・関数は複数の戻り値を返すことができる
・配列などに入れなくてもOK
・このような感じ → func 関数名(引数, 引数 型) (戻り値1の型, 戻り値2の型) {}
*/
