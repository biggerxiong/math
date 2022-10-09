package model

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type MidStream struct {
	ID int
	Point

	Cap    decimal.Decimal
	OriCap decimal.Decimal
}

func (s MidStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%.13f, Y:%.13f, Cap:%s}", s.ID, s.X, s.Y, s.Cap)
}

func (s MidStream) GetPoint() *Point {
	return &s.Point
}

func (s MidStream) GetID() int {
	return s.ID
}

func (s MidStream) GetX() float64 {
	return s.X
}

func (s MidStream) GetY() float64 {
	return s.Y
}
