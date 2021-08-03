package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sumをcに送る
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// 2つのgoroutine間で作業を分配し、両方のgoroutineで計算が完了すると最終結果が計算される。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // cから値を取得する

	fmt.Println(x, y, x+y)
}

/*
・チャンネルオペレータ
ch <- v    // v をチャネル ch へ送信する
v := <-ch  // ch から受信した変数を v へ割り当てる

・チャンネルを使う際の生成
ch := make(chan int)

片方が準備できるまでは送受信はブロックされるので、明確なロックや条件変数がなくても
goroutineは処理の同期を保つことができる。
*/
