## 5. 並列処理

### goroutine

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
goroutineはGoのランタイムに管理される軽量なスレッド。

goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要があります。
sync パッケージは同期する際に役に立つ方法を提供していますが、別の方法があるためそれほど必要ありません。
*/
```

### channel

```go
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
```

### closed channel

```go
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
```

### select

```go
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
```

### mutex

```go
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
```