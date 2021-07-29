package main

func sliceAppend() {
	var s []int // len=0 cap=0 []

	s = append(s, 0) // len=1 cap=1 [0]

	s = append(s, 1) // len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4) // len=5 cap=6 [0 1 2 3 4]
}

/*
もし、元の配列が、変数群を追加する際に
容量が小さい場合は、より大きいサイズの配列を割り当て直します。
その場合、戻り値となるスライスは、新しい割当先を示すようになります。
*/
