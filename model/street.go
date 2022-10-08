package main

// street 小区
type street struct {
	Point
}

func (p street) Key() string {
	return p.String()
}

var StreetToIndex map[string]int // 小区 -> 编号
