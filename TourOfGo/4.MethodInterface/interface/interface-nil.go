package main

import "fmt"

type III interface {
	M()
}

type TTT struct {
	S string
}

func (t *TTT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func mainnnn() {
	var i III

	var t *TTT
	i = t
	describee(i)
	i.M()

	i = &TTT{"Hello"}
	describee(i)
	i.M()
}

func describee(i III) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
インターフェース自体の中にある具体的な値がnilの場合、
メソッドはnilをレシーバーとして呼び出されます。

こういった挙動をエラーとする言語もあるが、Goでは
nilをレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的。

具体的な値としてnilを保持するインターフェースの値それ自体は非nilであることに注意。

ーーーー
また、ゼロ個のメソッドを指定されたインターフェース型は空のインターフェースと呼ばれる。
interface{}

空のインターフェースは任意の型の値を保持できる。
*/
