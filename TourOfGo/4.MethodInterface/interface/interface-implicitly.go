package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func mainn() {
	// 左側: interface 右側: なんらかの型
	// => ”右側の型に左側のinterfaceを実装する” と呼ぶ？
	var i I = T{"Hello"}
	i.M()
}

/*
インターフェースは暗黙的に実装される。

型にメソッドを実装していくことによって、インターフェースを実装する。
インターフェースを実装する際に明示的に宣言する必要はない(implementsのようなものは使わない)
*/
