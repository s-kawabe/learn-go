package main

type Point struct {
	X, Y int
}

func swap(p Point) {
	// swap X and Y
	x, y := p.X, p.Y
	p.X = x
	p.Y = y
}

p := Point{X: 1, Y: 2}
swap(p)

p.X // 1
p.Y // 2

// structは基本値型で処理される。
// 上の例ではswapにわたるPoint型の構造体はあくまでそのコピーが渡される為
// 元のPoint型構造体自体がswapされることはない。

// このような性質のため、メソッドなどではポインタレシーバにして
// 元の構造体が変更されるようにする必要がある。
// 以下は上記の問題を回避する例

func swap(p *Point) {
	// swap X and Y
	x, y := p.X, p.Y
	p.X = x
	p.Y = y
}

p := Point{X: 1, Y: 2}
swap(&p)
