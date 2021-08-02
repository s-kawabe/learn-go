package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main4() {
	v := Vertex2{3, 4}
	v.Scale2(5)
	// p := &v      これでも動く⭐️
	// p.Scale2(5)
	fmt.Println(v.Abs2())
	// p := *v
	// fmt.Println(p.Abs2())
}

/*
変数レシーバではScaleメソッドの操作は元のVertex変数のコピーを操作する。
main関数で宣言したVertex変数を変更するためにはScaleメソッドはポインタレシーバにする必要がある。

⭐️ ポインタレシーバの場合は呼び出し時に暗黙的にポインタに変換される
「v.Scale2(5)」のvはただの変数で、ポインタではないが、
Goが自動で「(&v).Scale(5)」という解釈を行う
*/
