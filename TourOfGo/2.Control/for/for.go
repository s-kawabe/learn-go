package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func moreFor() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func moreFor2() {
	for {
		// infinity loop
	}
}

/*
・初期化部分及び後処理ステートメントは省略することができる
・whileは無いので、moreForメソッドのような実装で同じような表現をする
・moreFor2メソッドは無限ループする。
*/
