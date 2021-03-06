package main

/*
sliceを指定する際のパターン
*/

//var a [10]int
// --- 上記の配列を示すslice(全て同じ) ---
// a[0:10]
// a[:10]
// a[0:]
// a[:]

func SlicePattern() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s = s[1:4] // [2, 3, 4]

	s = s[:2] // [2, 3]

	s = s[1:] // [3]
}
