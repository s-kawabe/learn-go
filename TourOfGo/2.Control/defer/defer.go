package main

import "fmt"

func main() {
	defer fmt.Println("Hello, defer!")

	fmt.Println("Hello!") // こちらが先に印字される
}

/*
deferへ渡した関数の実行を呼び出し元の関数の終わりまで遅延させる。
→jsのawaitとは少し違う？

しかし、deferへ渡した関数の引数はすぐに評価される。
*/

func loop() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/*
deferに渡した関数が複数ある場合、その呼び出しはスタックされる。
これらが実行されるとき(呼び出し元の関数のreturn時)LIFOの順序で実行される。
上記の例では数字の大きいものから少ないものへと印字される。
*/
