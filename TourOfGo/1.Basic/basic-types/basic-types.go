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
