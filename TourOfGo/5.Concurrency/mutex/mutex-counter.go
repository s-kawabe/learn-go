package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

/*
goroutineにおいて通信が必要なく、コンフリクトを避けるために常に
一度に1つのgoroutineだけが変数にアクセスできるようにしたい場合。
この手法を「排他制御（mutex）」と呼ぶ。

Goの標準ライブラリでは次の排他制御を提供している。
- sync.Mutex
	- Lock
	- Unlock

IncメソッドにあるようにLockとUnlockで囲むことで排他制御で実行するコードを定義できる。
ValueメソッドのようにmutexがUnlockされることを保証するためにdeferを使うこともできる。
*/
