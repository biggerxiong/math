package model

import "fmt"

type UpStream struct {
	ID int
	Point
	Cap int
}

func (s UpStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%s, Y:%s, Cap:%d}", s.ID, s.X, s.Y, s.Cap)
}

func (s UpStream) GetPoint() *Point {
	return &s.Point
}

func (s UpStream) GetID() int {
	return s.ID
}

func (s UpStream) GetX() string {
	return s.X
}

func (s UpStream) GetY() string {
	return s.Y
}
