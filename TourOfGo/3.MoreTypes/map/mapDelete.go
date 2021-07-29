package main

import "fmt"

func main2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // 42

	m["Answer"] = 57
	fmt.Println("The value:", m["Answer"]) // 57

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}

/*
キーに対する要素が存在するかどうか(okはbookean型)
elem, ok = m[key]

mapからキーを元に要素を削除する
delete(m, key)
*/
