package q2

// 以下の果物の価格の合計を出力するコードを書いてください
// m := map[string]int{
// 	"apple":  200,
// 	"banana": 300,
// 	"grapes": 150,
// 	"orange": 80,
// 	"papaya": 500,
// 	"kiwi":   90,
// }

func main() {

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	var result int
	for _, v := range m {
		result += v
	}
}
