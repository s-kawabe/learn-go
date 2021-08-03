# A Tour Of Go

## 1. 言語基礎

### 基礎の型

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
変数宣言はまとめることができる。
特別な理由がない限り整数の変数はuintではなくintを使用する。
*/

/* Go言語における基本の組み込み型
bool

string

int  int8  int16  int32  int64 // 符号あり(マイナス値を含んだ範囲)
uint uint8 uint16 uint32 uint64 uintptr // 符号なし(マイナス値を含まない範囲)

byte // uint8 の別名

rune // int32 の別名
	// Unicode のコードポイントを表す

float32 float64

complex64 complex128
*/

```

### 定数

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
・定数が使えるのは文字(character)、文字列(string)、boolean、数値(numeric)のみ
・定数は := を使って宣言することはできない
・定数はパスカルケース？
*/

```

### 変数

```go
package main

import (
	"fmt"
)

var c, python, java bool
var cpp, csharp, javascript = true, true, true

func main() {
	var i int
	k := 3 // short variable declaration

	fmt.Println(i, c, python, java)      // 0 false false false
	fmt.Println(cpp, csharp, javascript) // 0 true true true
	fmt.Println(k)                       // 3
}

/*
・varで変数を宣言する。パッケージグローバルまたは関数内で使用可能。
・宣言と同時に初期化子を与える場合は型の指定は省略できる。
・初期化していない物を使用してもエラーにならず、仕様による初期値が与えられる⭐️ → zero values
・「:=」を使うと値の指定のみで、暗黙的な初期化と型定義が行われる。
　これは関数外では使用できない。
*/

/* zero valuesの仕様
-----
数値型(int,floatなど): 0
bool型: false
string型: "" (空文字列( empty string ))
-----
*/
```

### 型変換

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

/*
変数 v 、型 T があった場合、 T(v) は、変数 v を T 型へ変換します。
Goにおいては暗黙的な型変換は行われず、常に明示的な指定が必要となる。
*/
```

### 関数

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 20))
}

/*
関数は、0個以上の引数を取ることができます。
型は変数の後に記述する。
*/

func add2(x, y int) int {
	return x + y
}

/* 型の省略記法
型の2つ以上の引数が同じ型である場合には上記のような記述ができる
*/
```

### 複数の戻り値

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/*
・関数は複数の戻り値を返すことができる
・配列などに入れなくてもOK
・このような感じ → func 関数名(引数, 引数 型) (戻り値1の型, 戻り値2の型) {}
*/
```

### 関数の名前付き戻り値

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	// 既にxとyは初期化されている
	x = sum * 4 / 9
	y = sum - x
	return // xとyを返す
}

func main() {
	fmt.Println(split(17))
}

/*
戻り値となる変数に名前をつけることができる。
これを指定すると、return　のみでも記載した変数を返すことができる。 → (naked return)
指定した変数は、関数内の最初に定義されたと同等になる。

→長い関数で使用すると可読性が下がる！！
*/
```

### エクスポート

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) → undifined
	fmt.Println(math.Pi)
}

/*
Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できる
エクスポートされた名前である。
例えば Pi はmathパッケージでエクスポートされている。

小文字で始まるpiなどはエクスポートされていない名前。
パッケージをimportすると、そのパッケージがエクスポートしている名前を
参照することができる。 エクスポートされていない名前(小文字で始まる名前)は
直接外部パッケージからアクセスすることができない
*/
```

### パッケージ

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
規約で、パッケージ名はインポートパスの最後の要素と同じ名前になります。
例えば、インポートパスが "math/rand" のパッケージは、
package rand ステートメントで始まるファイル群で構成します。
*/
```

---

## 2. 制御構造

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

---

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

---

## 4. メソッドとインターフェース

### メソッド

#### 値レシーバメソッド

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// (v Vertex) vはレシーバ引数 Vertexはレシーバ型
// ポインタではないstructを使う → 値レシーバ
// このvはあくまでVertexのコピーになる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// レシーバ引数(メソッド)を使わず関数で記述した場合
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main2() {
	v := Vertex{3, 4}
	fmt.Println(Abs2(v))
}

/*
・メソッドは「レシーバ引数を伴う関数」
・Goにはクラスはないが、型にメソッドを定義できる。
・関数名の前にレシーバ引数を指定できる。
・レシーバは、func キーワードとメソッド名の間に自身の引数リストで表現します。
*/

/*
関数定義は func (レシーバ値 レシーバ型) 関数名という形になる。
特に*Tで定義されたレシーバをポインタレシーバ、Tを値レシーバと呼ぶ。
*/
```

#### ポインタレシーバメソッド

```go
package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main4() {
	v := Vertex2{3, 4}
	v.Scale2(5)
	// p := &v      これでも動く⭐️
	// p.Scale2(5)
	fmt.Println(v.Abs2())
	// p := *v
	// fmt.Println(p.Abs2())
}

/*
変数レシーバではScaleメソッドの操作は元のVertex変数のコピーを操作する。
main関数で宣言したVertex変数を変更するためにはScaleメソッドはポインタレシーバにする必要がある。

⭐️ ポインタレシーバの場合は呼び出し時に暗黙的にポインタに変換される
「v.Scale2(5)」のvはただの変数で、ポインタではないが、
Goが自動で「(&v).Scale(5)」という解釈を行う
*/
```

#### ポインタレシーバを知る

```go
package main

/*
⭐️ ポインターレシーバを使う理由
・メソッドがレシーバが刺す先の変数を変更するため
・メソッドの呼び出し毎に変数のコピーを避けるため(例えばレシーバが大きな構造体である場合)

基本的に値レシーバかポインタレシーバどちらかで揃える。
*/

import (
	"fmt"
	"math"
)

type Vertexxx struct {
	X, Y float64
}

func (v *Vertexxx) Scaleee(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertexxx) Absee() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func mainnn() {
	v := Vertexxx{3, 4} // &Vertexxx{3, 4} でもOK
	fmt.Printf("Before scaling: %+v, Absee: %v\n", v, v.Absee())
	v.Scaleee(5)
	fmt.Printf("After scaling: %+v, Absee: %v\n", v, v.Absee())
}
```

#### struct以外のメソッド定義

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// ここで(f float64)と直接組み込み型を指定することはできない
// ⭐️ パッケージ内でtype定義された型のみがレシーバにできる
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

/*
struct以外の型でもメソッドを宣言できる。
恐らくレシーバに指定したものに向けてそのメソッドがぶら下がるイメージ。

レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要があります。
他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できません
*/
```

### インターフェース

#### インターフェース

```go
package main

import (
	"fmt"
	"math"
)

// interfaceの中にはメソッドのリストを記載する
type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloatはAbserを実装している?
	a = &v // *VertexはAbserを実装している?

	// a = v
	// Absメソッドが*Vertexの定義であり、VertexがAbserインターフェース
	// を実装していないのでエラーになる。

	fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
・interface型はメソッドのシグネチャの集まりで定義する。
・メソッドの集まりを実装した値を、interface型の変数へ持たせることができる。
*/
```

#### インターフェースの値

```go
package main

import (
	"fmt"
	"math"
)

type II interface {
	M()
}

type TT struct {
	S string
}

func (t *TT) M() {
	println(t.S)
}

type FF float64

// 同じ名前のメソッドが2個あっても、それがどこにぶら下がっているかを
// Goは適切に認識してくれるためこれでOK
func (f FF) M() {
	fmt.Println(f)
}

func mainnn() {
	var i I

	i = &TT{"Hello"}
	describe(i)
	i.M() // ここのM()と

	i = FF(math.Pi)
	describe(i)
	i.M() // ここのM()を適切に分けて認識している
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
(value, type)
のように、インターフェースの値は、値と具体的な型のタプルのように考えることができる。

インターフェースの値は、特定の基底になる具体的な型の値を保持します。
インターフェースの値のメソッドを呼び出すと、その基底型の同じ名前のメソッドが実行される。
*/
```

#### 暗黙的なインターフェース実装

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func mainn() {
	// 左側: interface 右側: なんらかの型
	// => ”右側の型に左側のinterfaceを実装する” と呼ぶ？
	var i I = T{"Hello"}
	i.M()
}

/*
インターフェースは暗黙的に実装される。

型にメソッドを実装していくことによって、インターフェースを実装する。
インターフェースを実装する際に明示的に宣言する必要はない(implementsのようなものは使わない)
*/
```

#### nilインターフェース

```go
package main

import "fmt"

type III interface {
	M()
}

type TTT struct {
	S string
}

func (t *TTT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func mainnnn() {
	var i III

	var t *TTT
	i = t
	describee(i)
	i.M()

	i = &TTT{"Hello"}
	describee(i)
	i.M()
}

func describee(i III) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
インターフェース自体の中にある具体的な値がnilの場合、
メソッドはnilをレシーバーとして呼び出されます。

こういった挙動をエラーとする言語もあるが、Goでは
nilをレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的。

具体的な値としてnilを保持するインターフェースの値それ自体は非nilであることに注意。

ーーーー
また、ゼロ個のメソッドを指定されたインターフェース型は空のインターフェースと呼ばれる。
interface{}

空のインターフェースは任意の型の値を保持できる。
*/
```

#### stringerインターフェースとは

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// コンストラクター的な感じで使える
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	// => Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}

/*
fmtパッケージには以下のようなinterfaceが定義されている
	type Stringer interface {
    	String() string
	}

これをオーバーライドするような形にもすることができる。
*/
```

### 型アサーション

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic
	fmt.Println(f)
}

/* ⭐️ 型アサーション
    t := i.(T)
インターフェースの値iが具体的な型Tを保持し、基になる
Tの値を変数tに代入することを表している。
iがTを保持していない場合はpanicを起こす。

	t, ok := i.(T)
このようにするとiがTを保持していない場合でもpanicを起こさない。
- iがTを保持している場合
→tはアサーション後の値になり、okはtrueになる
- iがTを保持していない場合
→tはゼロ値になり okはfalseになる。
*/
```

### 型switch

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main() {
	do(21)      // Twice 21 is 42
	do("hello") // "hello" is 5 bytes long
	do(true)    // I don't know about type bool!
}

/* ⭐️ type switch
型switchのcaseは型を指定し、それらの値は指定された
インターフェースの値が保持する値の型と比較される。

型switchのswitchの部分に指定する宣言は型アサーションと同じ構文を持つが
typeはcase文のそれぞれの型に置き換えられる。
*/
```

### エラー処理

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// error型はfmt.Stringerに似た、組み込みのインターフェース。
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/*
fmtパッケージは変数を文字列で出力する際にerrorインターフェースを確認する。
type error interface {
    Error() string
}
*/
```

### reader

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") // string型のデータストリーム?
	b := make([]byte, 8)                     // バイトスライス
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
ioパッケージにはデータストリームを読むことを表現する「io.Reader」インターフェースを持っている。
このインターフェースからの派生にはファイル、ネットワーク接続、圧縮、暗号化など多くの実装がある。

io.ReaderインターフェースはReadメソッドを持つ。
func (T) Read(b []byte) (n int, err error)

Readはデータを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返す。
io.EOF のエラーを返した場合はストリームの終端であることを表す。

上記はstrings.Readerを作成し、8byte毎に読み出している。
*/
```

---

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