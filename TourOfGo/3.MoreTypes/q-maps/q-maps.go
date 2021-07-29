// WordCount 関数を実装してみましょう。
// string s で渡される文章の、各単語の出現回数のmapを返す必要があります。
// wc.Test 関数は、引数に渡した関数に対しテストスイートを実行し、
// 成功か失敗かを結果に表示します。

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splited := strings.Fields(s)
	var m = make(map[string]int)
	for _, v := range splited {
		v2, ok := m[v] // ここがとても冗長、m[v]++ でOK
		if ok {
			m[v] = v2 + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

// 模範解答
// func WordCount(s string) map[string]int {
//     m := make(map[string]int)
//     for _, w := range strings.Fields(s) {
//         m[w]++
//     }
//     return m
// }

func main() {
	wc.Test(WordCount)
}
