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
