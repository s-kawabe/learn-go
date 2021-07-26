package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.pi) → undifined
	fmt.Println(math.Pi)
}

/*
Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できる
エクスポートされた名前である。
例えば Pi はmathパッケージでエクスポートされている。

小文字で始まるpiなどはエクスポートされていない名前。
パッケージをimportすると、そのパッケージがエクスポートしている名前を
参照することができる。 エクスポートされていない名前(小文字で始まる名前)は
直接外部パッケージからアクセスすることができない
*/
