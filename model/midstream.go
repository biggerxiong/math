package model

import "fmt"

type MidStream struct {
	ID  int
	X   string
	Y   string
	Cap int
}

func (s MidStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%s, Y:%s, Cap:%d}", s.ID, s.X, s.Y, s.Cap)
}
