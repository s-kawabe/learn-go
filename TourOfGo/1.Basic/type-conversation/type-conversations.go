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
