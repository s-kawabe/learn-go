package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

/*
規約で、パッケージ名はインポートパスの最後の要素と同じ名前になります。
例えば、インポートパスが "math/rand" のパッケージは、
package rand ステートメントで始まるファイル群で構成します。
*/
