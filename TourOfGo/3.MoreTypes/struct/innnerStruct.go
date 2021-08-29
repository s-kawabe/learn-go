package main

type Feed struct {
	Name string
	Amount uint
}

type Animal struct {
	Name string
	Feeds Feed // Feed型の埋め込み Feed のみの記述でもOK
}

a := Animal{
	Name: "Tiger",
	Feed: Feed struct {
		Name: "Meat",
		Amount: 10
	}
}

a.Name //= "Tiger"
a.Feed.Name //= "Meat"
a.Feed.Amount //= 10

/*
	複合リテラルの内側にさらに複合リテラルを書くことができる。(その際のフィールド名の値の名前は同じで良い)
	structの中のstructの中のフィールドという感じで参照する際は a.Feed.Amount などと記述する
	この時、フィールド名が一意に決まるならば中間のstructを省略して記述できる
*/