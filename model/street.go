package model

import "fmt"

// Street 小区
type Street struct {
	ID            int
	BuildingCount int
	FamilyCount   int
	PeopleCount   int
	X             string
	Y             string
	StreetIndex   string
	BelongTo      string
}

func (s Street) String() string {
	return fmt.Sprintf("Street{ID:%d, X:%s, Y:%s}", s.ID, s.X, s.Y)
}

func (s Street) Key() string {
	return s.String()
}

var StreetToIndex map[string]int // 小区 -> 编号
