package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // xをチャンネルcに送れるなら
			x, y = y, x+y
		case <-quit: // quitチャンネルから受信ができるなら
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// 以下の即時実行関数は、実行完了を待たずに次の行に進む？
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

/*
selectステートメントはgoroutineを複数の通信操作で待たせる。
selectは複数あるcaseのいずれかが準備できるようになるまでブロックし、
準備ができたcaseを実行する。
もし複数のcaseが準備できた場合caseはランダムに選択される。

どのcaseも準備ができていなければselectのなかのdefaultが実行される。
*/
