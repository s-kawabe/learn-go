package main

import "fmt"

// intを返す関数を返す関数
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// intを返す関数を返す関数が入った変数
	f := fibonacci()
	for i := 0; i < 10; i++ {
		// intを返す関数を返す関数が入った変数を関数として実行
		fmt.Println(f())
	}
}
