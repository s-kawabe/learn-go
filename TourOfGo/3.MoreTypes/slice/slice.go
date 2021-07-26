package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // => [3 5 7]

	SliceChange()
}

/*
配列が固定長なのに対し、スライスは可変長である。
配列よりもスライスの方が一般的。
primes[1:4] のようのコロンで区切られた境界を示すことでスライスが作られる。
*/

func SliceChange() {
	names := [4]string{
		"Alice",
		"Bob",
		"Charlie",
		"David",
	}
	fmt.Println(names) // => [Alice Bob Charlie David]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // => [Alice Bob] [Bob Charlie]

	b[0] = "Eve"
	fmt.Println(a, b)  // => [Alice Eve] [Eve Charlie]
	fmt.Println(names) // => [Alice Eve Charlie David]
}

/*
スライスは配列への参照のようなもの。
スライスの要素を変更すると、その元となる配列の対応する要素が変更される。
*/
