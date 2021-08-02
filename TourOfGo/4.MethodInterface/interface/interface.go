package main

import (
	"fmt"
	"math"
)

// interfaceの中にはメソッドのリストを記載する
type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloatはAbserを実装している?
	a = &v // *VertexはAbserを実装している?

	// a = v
	// Absメソッドが*Vertexの定義であり、VertexがAbserインターフェース
	// を実装していないのでエラーになる。

	fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
・interface型はメソッドのシグネチャの集まりで定義する。
・メソッドの集まりを実装した値を、interface型の変数へ持たせることができる。
*/
