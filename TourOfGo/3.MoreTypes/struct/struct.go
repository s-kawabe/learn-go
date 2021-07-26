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
