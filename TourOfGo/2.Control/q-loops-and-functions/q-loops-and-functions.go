package main

import (
	"fmt"
	"math"
)

func MySqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(MySqrt(2))
	fmt.Println(math.Sqrt(2))
}

// Sqrt → 平方根を求める
// 平方根→ aを2乗するとbになるとき、 aはbの平方根。(4(b)の平方根は2(a)である。)
