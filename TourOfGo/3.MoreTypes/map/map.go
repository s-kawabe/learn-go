package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // make関数によるmap作成
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["aaa"] = Vertex{
		1.68433, -2.39967,
	}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} aaa:{1.68433 -2.39967}]
}

func mapLiteral() {
	var mm = map[string]Vertex{ // リテラルによるmap作成
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mm) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// 上の省略ver
	// mapに渡すトップレベルの型が単純な場合は推論できるため省略可能
	var mmm = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(mmm) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

/*
・mapには、キーと値のペアが格納されている。
・mapのzero valuesはnilである。
・make関数は指定された型のmapを初期化して、使用可能な状態で返す。
    → make(map[キーの型]値の型, キャパシティの初期値)
・または初期値を指定して初期化する
    → var 変数名 map[key]value = map[key]value{key1: value1, key2: value2}
*/
