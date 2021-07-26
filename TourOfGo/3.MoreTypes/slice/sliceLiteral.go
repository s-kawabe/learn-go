package main

import "fmt"

func literal() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// 配列リテラルの例
//  [3]bool{true, false, true}

// 以下は上記と同様の配列を作成した後、それを参照するスライスを作成している
//  []bool{true, false, true}
