package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // 1個の別スレッドgoroutineが動くことを伝える
	// これをやったらwg.Done()を必ずやる必要がある
	go goroutine("huga", &wg)
	normal("hoge")
	wg.Wait() // wg.Doneされるまで待つ
}

// =>
// huga
// hoge
// huga
// hoge
// hoge
// huga
// huga
// hoge
// hoge
// huga

// sleepを入れないとgoroutine化したスレッドが終わるよりも先にメインスレッドが終了してしまうことがある
// => sync.waitGroupが使える
