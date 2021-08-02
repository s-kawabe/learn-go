package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	f = i.(float64) // panic
	fmt.Println(f)
}

/* ⭐️ 型アサーション
    t := i.(T)
インターフェースの値iが具体的な型Tを保持し、基になる
Tの値を変数tに代入することを表している。
iがTを保持していない場合はpanicを起こす。

	t, ok := i.(T)
このようにするとiがTを保持していない場合でもpanicを起こさない。
- iがTを保持している場合
→tはアサーション後の値になり、okはtrueになる
- iがTを保持していない場合
→tはゼロ値になり okはfalseになる。
*/
