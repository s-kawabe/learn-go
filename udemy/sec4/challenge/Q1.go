package q1

// 以下のスライスから一番小さい数を探して出力するコードを書いてください。
// l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

func main() {
  result := 10000
	for _, v := range l {
		if(v < result) {
			result = v
		}
	}
	fmt.Println("most minimum number is %d", result)
}
