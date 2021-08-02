package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

func main() {
	do(21)      // Twice 21 is 42
	do("hello") // "hello" is 5 bytes long
	do(true)    // I don't know about type bool!
}

/* ⭐️ type switch
型switchのcaseは型を指定し、それらの値は指定された
インターフェースの値が保持する値の型と比較される。

型switchのswitchの部分に指定する宣言は型アサーションと同じ構文を持つが
typeはcase文のそれぞれの型に置き換えられる。
*/
