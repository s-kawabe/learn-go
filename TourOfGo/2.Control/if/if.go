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
