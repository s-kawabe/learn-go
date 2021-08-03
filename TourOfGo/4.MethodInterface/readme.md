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