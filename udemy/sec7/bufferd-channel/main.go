package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

// =>
// 1
// 2
// 100
// 200

// channelのバッファを決める
// rangeでループする場合はchannelをcloseする必要がある
