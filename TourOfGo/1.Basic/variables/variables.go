package main

import (
	"fmt"
)

var c, python, java bool
var cpp, csharp, javascript = true, true, true

func main() {
	var i int
	k := 3 // short variable declaration

	fmt.Println(i, c, python, java)      // 0 false false false
	fmt.Println(cpp, csharp, javascript) // 0 true true true
	fmt.Println(k)                       // 3
}

/*
・varで変数を宣言する。パッケージグローバルまたは関数内で使用可能。
・宣言と同時に初期化子を与える場合は型の指定は省略できる。
・初期化していない物を使用してもエラーにならず、仕様による初期値が与えられる⭐️ → zero values
・「:=」を使うと値の指定のみで、暗黙的な初期化と型定義が行われる。
　これは関数外では使用できない。
*/

/* zero valuesの仕様
-----
数値型(int,floatなど): 0
bool型: false
string型: "" (空文字列( empty string ))
-----
*/
