package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// error型はfmt.Stringerに似た、組み込みのインターフェース。
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/*
fmtパッケージは変数を文字列で出力する際にerrorインターフェースを確認する。
type error interface {
    Error() string
}
*/
