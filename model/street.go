package model

import "fmt"

// Street 小区
type Street struct {
	ID int
	Point

	BuildingCount int
	FamilyCount   int
	PeopleCount   int
	StreetIndex   string
	BelongTo      string

	Cap string
}

func (s Street) String() string {
	return fmt.Sprintf("Streets{ID:%d, X:%s, Y:%s}", s.ID, s.X, s.Y)
}

func (s Street) Key() string {
	return s.String()
}

func (s Street) GetPoint() *Point {
	return &s.Point
}

func (s Street) GetID() int {
	return s.ID
}

func (s Street) GetX() string {
	return s.X
}

func (s Street) GetY() string {
	return s.Y
}
