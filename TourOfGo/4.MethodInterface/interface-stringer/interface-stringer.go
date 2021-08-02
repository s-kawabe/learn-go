package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// コンストラクター的な感じで使える
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	// => Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
}

/*
fmtパッケージには以下のようなinterfaceが定義されている
	type Stringer interface {
    	String() string
	}

これをオーバーライドするような形にもすることができる。
*/
