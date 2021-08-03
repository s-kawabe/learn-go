package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
・チャンネルがこれ以上送信する値がないことを示すためのclose
v, ok := <-ch
受診する値がなくチャンネルが閉じているならokはfalseになる

・チャンネルが閉じられるまでチャンネルからの値を受信し続ける
for i := range ch

※ closeするのは送り手のチャネルだけにし、受け手はcloseしてはいけない
closeしたチャネルに送信をするとpanicが発生する

※ チャンネルはファイルと違って通常closeする必要はない。
*/
