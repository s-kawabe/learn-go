package main

import (
	"fmt"
	"math"
)

type II interface {
	M()
}

type TT struct {
	S string
}

func (t *TT) M() {
	println(t.S)
}

type FF float64

// 同じ名前のメソッドが2個あっても、それがどこにぶら下がっているかを
// Goは適切に認識してくれるためこれでOK
func (f FF) M() {
	fmt.Println(f)
}

func mainnn() {
	var i I

	i = &TT{"Hello"}
	describe(i)
	i.M() // ここのM()と

	i = FF(math.Pi)
	describe(i)
	i.M() // ここのM()を適切に分けて認識している
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
(value, type)
のように、インターフェースの値は、値と具体的な型のタプルのように考えることができる。

インターフェースの値は、特定の基底になる具体的な型の値を保持します。
インターフェースの値のメソッドを呼び出すと、その基底型の同じ名前のメソッドが実行される。
*/
