## 言語基礎

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