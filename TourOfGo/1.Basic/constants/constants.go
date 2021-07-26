package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
・定数が使えるのは文字(character)、文字列(string)、boolean、数値(numeric)のみ
・定数は := を使って宣言することはできない
・定数はパスカルケース？
*/
