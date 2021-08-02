package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// (v Vertex) vはレシーバ引数 Vertexはレシーバ型
// ポインタではないstructを使う → 値レシーバ
// このvはあくまでVertexのコピーになる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// レシーバ引数(メソッド)を使わず関数で記述した場合
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main2() {
	v := Vertex{3, 4}
	fmt.Println(Abs2(v))
}

/*
・メソッドは「レシーバ引数を伴う関数」
・Goにはクラスはないが、型にメソッドを定義できる。
・関数名の前にレシーバ引数を指定できる。
・レシーバは、func キーワードとメソッド名の間に自身の引数リストで表現します。
*/

/*
関数定義は func (レシーバ値 レシーバ型) 関数名という形になる。
特に*Tで定義されたレシーバをポインタレシーバ、Tを値レシーバと呼ぶ。
*/
