package main

import "fmt"

func goroutine(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c) // CHANNELが終了することを明示
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string, 4)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

// => 以下の出力になるように直す
// test1!
// test1!test2!
// test1!test2!test3!
// test1!test2!test3!test4!
