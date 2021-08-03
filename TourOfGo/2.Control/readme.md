## 制御構造

### for

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func moreFor() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func moreFor2() {
	for {
		// infinity loop
	}
}

/*
・初期化部分及び後処理ステートメントは省略することができる
・whileは無いので、moreForメソッドのような実装で同じような表現をする
・moreFor2メソッドは無限ループする。
*/
```

### if

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // ⭐️
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
・for同様にifでも()はつけない
*/

// if short statement
func pow(x, m, lim float64) float64 {
	if v := math.Pow(x, m); v < lim {
		return v
	}
	return lim
}

/*
このような書き方をするとそのifブロック内でのみ有効な変数を定義できる。
→elseを定義した場合はその中でも使える。
*/
```

### switch

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// switchの条件省略 (switch trueと同じ)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

/*
Goのswitchはreturnを書かずとも自動的にbreakされ、次のcaseに流れることはない。
*/
```