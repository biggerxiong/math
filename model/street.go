package model

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Street 小区
type Street struct {
	ID int
	Point

	BuildingCount int
	FamilyCount   int
	PeopleCount   int
	StreetIndex   string
	BelongTo      string

	Cap decimal.Decimal
}

func (s Street) String() string {
	return fmt.Sprintf("Streets{ID:%d, X:%.13f, Y:%.13f}", s.ID, s.X, s.Y)
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

func (s Street) GetX() float64 {
	return s.X
}

func (s Street) GetY() float64 {
	return s.Y
}
