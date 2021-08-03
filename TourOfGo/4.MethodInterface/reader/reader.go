package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") // string型のデータストリーム?
	b := make([]byte, 8)                     // バイトスライス
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

/*
ioパッケージにはデータストリームを読むことを表現する「io.Reader」インターフェースを持っている。
このインターフェースからの派生にはファイル、ネットワーク接続、圧縮、暗号化など多くの実装がある。

io.ReaderインターフェースはReadメソッドを持つ。
func (T) Read(b []byte) (n int, err error)

Readはデータを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返す。
io.EOF のエラーを返した場合はストリームの終端であることを表す。

上記はstrings.Readerを作成し、8byte毎に読み出している。
*/
