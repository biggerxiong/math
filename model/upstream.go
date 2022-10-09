package model

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type UpStream struct {
	ID int
	Point
	Cap decimal.Decimal
}

func (s UpStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%.13f, Y:%.13f, Cap:%s}", s.ID, s.X, s.Y, s.Cap)
}

func (s UpStream) GetPoint() *Point {
	return &s.Point
}

func (s UpStream) GetID() int {
	return s.ID
}

func (s UpStream) GetX() float64 {
	return s.X
}

func (s UpStream) GetY() float64 {
	return s.Y
}
