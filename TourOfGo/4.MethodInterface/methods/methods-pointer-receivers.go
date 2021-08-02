package main

/*
⭐️ ポインターレシーバを使う理由
・メソッドがレシーバが刺す先の変数を変更するため
・メソッドの呼び出し毎に変数のコピーを避けるため(例えばレシーバが大きな構造体である場合)

基本的に値レシーバかポインタレシーバどちらかで揃える。
*/

import (
	"fmt"
	"math"
)

type Vertexxx struct {
	X, Y float64
}

func (v *Vertexxx) Scaleee(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertexxx) Absee() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func mainnn() {
	v := Vertexxx{3, 4} // &Vertexxx{3, 4} でもOK
	fmt.Printf("Before scaling: %+v, Absee: %v\n", v, v.Absee())
	v.Scaleee(5)
	fmt.Printf("After scaling: %+v, Absee: %v\n", v, v.Absee())
}
