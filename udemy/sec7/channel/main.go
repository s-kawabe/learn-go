package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // chanはqueueのような構造になっている
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c // goroutine1でsumが受信されるまで同期的に待つ
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}

// channnelは1つでもいいし、複数作って別のデータを入れることもできる。
// 要件によって切り替える
