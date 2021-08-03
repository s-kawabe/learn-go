## 3. より厳密な型

### 値としての関数

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

### 配列

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // => Hello World
	fmt.Println(a)          // [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // => [2 3 5 7 11 13]
}

/*
配列の長さは型の一部分である。
Goにおける"配列"はサイズを変えることはできない。 => 固定長
*/
```

### スライス

#### スライス基礎

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // => [3 5 7]

	SliceChange()
}

/*
配列が固定長なのに対し、スライスは可変長である。
配列よりもスライスの方が一般的。
primes[1:4] のようのコロンで区切られた境界を示すことでスライスが作られる。
*/

func SliceChange() {
	names := [4]string{
		"Alice",
		"Bob",
		"Charlie",
		"David",
	}
	fmt.Println(names) // => [Alice Bob Charlie David]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // => [Alice Bob] [Bob Charlie]

	b[0] = "Eve"
	fmt.Println(a, b)  // => [Alice Eve] [Eve Charlie]
	fmt.Println(names) // => [Alice Eve Charlie David]
}

/*
スライスは配列への参照のようなもの。
スライスの要素を変更すると、その元となる配列の対応する要素が変更される。
*/
```

#### スライスのlenとcap

```go
package main

import "fmt"

func lenAndCap() {
	s := []int{1, 2, 3, 4, 5}
	printSlice(s) // => len=5 cap=5 [1 2 3 4 5]

	s = s[:0]
	printSlice(s) // => len=0 cap=5 []

	s = s[:4]
	printSlice(s) // => len=4 cap=5 [1 2 3 4]

	s = s[2:]
	printSlice(s) // => len=2 cap=5 [3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
スライスのzero valuesはnil。
var s []int // => sはnil

nilスライスは0の長さと容量を持っており、何の元となる配列も持っていません。
*/
```

#### スライスへの追加

```go
package main

func sliceAppend() {
	var s []int // len=0 cap=0 []

	s = append(s, 0) // len=1 cap=1 [0]

	s = append(s, 1) // len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4) // len=5 cap=6 [0 1 2 3 4]
}

/*
もし、元の配列が、変数群を追加する際に
容量が小さい場合は、より大きいサイズの配列を割り当て直します。
その場合、戻り値となるスライスは、新しい割当先を示すようになります。
*/
```

#### スライスリテラル

```go
package main

import "fmt"

func literal() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// 配列リテラルの例
//  [3]bool{true, false, true}

// 以下は上記と同様の配列を作成した後、それを参照するスライスを作成している
//  []bool{true, false, true}
```

#### スライスの作成

```go
package main

import (
	"fmt"
	"strings"
)

func main2() {
	a := make([]int, 5) // nilではなく0で埋められたスライスを生成 ⭐️
	// len=5 cap=5 [0 0 0 0 0]
	printSlice2("a", a)

	//        型    長さ 最大要素数
	b := make([]int, 0, 5)
	printSlice2("b", b) // len=5 cap=5 []

	c := b[:2]          // capは変わらない
	printSlice2("c", c) // len=2 cap=5 [0 0]

	d := c[2:5]         // cap変わる
	printSlice2("d", d) // len=3 cap=3 [0 0 0]
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// sliceの中にsliceを入れることができる
func twoSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

#### スライスのパターン

```go
package main

/*
sliceを指定する際のパターン
*/

//var a [10]int
// --- 上記の配列を示すslice(全て同じ) ---
// a[0:10]
// a[:10]
// a[0:]
// a[:]

func SlicePattern() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s = s[1:4] // [2, 3, 4]

	s = s[:2] // [2, 3]

	s = s[1:] // [3]
}
```

### range

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
rangeは、スライスやマップをを一つずつ反復処理するために使用する。
スライスをrangeに使用する際の2つの変数
1. インデックス
2. インデックスの場所の要素のコピー
*/

func dust() {
	pow := make([]int, 10)
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

/*
インデックスや値は "_" を指定すると捨てられる
値がいらずインデックスのみ欲しい場合は普通に値を書かなければ良い
*/
```

### map

#### mapの作成

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // make関数によるmap作成
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["aaa"] = Vertex{
		1.68433, -2.39967,
	}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} aaa:{1.68433 -2.39967}]
}

func mapLiteral() {
	var mm = map[string]Vertex{ // リテラルによるmap作成
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mm) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// 上の省略ver
	// mapに渡すトップレベルの型が単純な場合は推論できるため省略可能
	var mmm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(mmm) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

/*
・mapには、キーと値のペアが格納されている。
・mapのzero valuesはnilである。
・make関数は指定された型のmapを初期化して、使用可能な状態で返す。
    → make(map[キーの型]値の型, キャパシティの初期値)
・または初期値を指定して初期化する
    → var 変数名 map[key]value = map[key]value{key1: value1, key2: value2}
*/
```

#### mapの要素削除

```go
package main

import "fmt"

func main2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // 42

	m["Answer"] = 57
	fmt.Println("The value:", m["Answer"]) // 57

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}

/*
キーに対する要素が存在するかどうか(okはboolean型)
elem, ok = m[key]

mapからキーを元に要素を削除する
delete(m, key)
*/
```

### ポインター

```go
package main

import "fmt"

// Goではポインタを扱う。 変数 T のポインタは *T 型。
// var p *int

// &オペレータは、値をもとにポインタを引き出す。
// i := 42
// p = &i

// *オペレータは、ポインタの指す先の変数を示す。
// fmt.Println(*p)
// *p = 21 =>これをやると元のiも変わる

func main() {
	i, j := 42, 2701

	p := &i         // pはiのポインタ
	fmt.Println(*p) // *pでポインタから値を割り出す => 42
	*p = 21         // i = 21と同じことになる? ポインタに直接書き込んでいる?
	fmt.Println(i)  // => 21

	p = &j // pにjのポインタを入れる
	*p = *p / 37
	fmt.Println(j)
}

// *オペレータを使うと使用した変数の元の値にまで影響が及ぶ？
// 値のポインタと型のポインタがある
```

### 構造体 struct

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // => {1 2}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // => 4

	p := &v
	p.X = 10
	fmt.Println(v) // => {10 2}
}

/*
・struct(構造体)は、フィールドの集まりである
・structを初期化した後。「変数.値」としてアクセスする
・structのポインタを通してアクセスすることもできる。
*/

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1} // Yは0
	v3 = Vertex2{}     // XもYも0
	p  = &Vertex2{1, 2}
)

func main2() {
	fmt.Println(v1, p, v2, v3) // => {1 2} &{1 2} {1 0} {0 0}
}
```