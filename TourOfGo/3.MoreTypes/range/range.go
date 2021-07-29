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
