package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
チャンネルはバッファとして使用できる。

・バッファを持つチャンネルの初期化
ch := make(chan int, 100)

バッファが詰まった際はチャンネルへの送信をブロックする。
また、バッファが空の時には、チャンネルの受診をブロックする。
*/
