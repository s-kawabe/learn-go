package main

type Point struct{ X, Y int }

ps := make{[]Point, 5}
for _, p := range ps {
	fmt.Println(p.X, p.Y)
}

/* 結果
0 0
0 0
0 0
0 0
0 0
*/

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return str
}

ps := Points{}
ps = append(ps, &Point{1,2})
ps = append(ps, nil)
ps = append(ps, &Point{3,4})
ps.ToString()

// []*Pointのように複雑な宣言に対してtypeによるエイリアスを定義してそのエイリアスに対して
// メソッドを定義することで型を扱いやすくする