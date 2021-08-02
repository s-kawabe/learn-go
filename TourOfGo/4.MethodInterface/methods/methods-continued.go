package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// ここで(f float64)と直接組み込み型を指定することはできない
// ⭐️ パッケージ内でtype定義された型のみがレシーバにできる
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
struct以外の型でもメソッドを宣言できる。
恐らくレシーバに指定したものに向けてそのメソッドがぶら下がるイメージ。

レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。
他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できません
*/
