package model

import "fmt"

type MidStream struct {
	ID int
	Point

	Cap int
}

func (s MidStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%s, Y:%s, Cap:%d}", s.ID, s.X, s.Y, s.Cap)
}

func (s MidStream) GetPoint() *Point {
	return &s.Point
}

func (s MidStream) GetID() int {
	return s.ID
}

func (s MidStream) GetX() string {
	return s.X
}

func (s MidStream) GetY() string {
	return s.Y
}
