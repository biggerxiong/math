package model

import "fmt"

type UpStream struct {
	ID  int
	X   string
	Y   string
	Cap int
}

func (s UpStream) String() string {
	return fmt.Sprintf("MidStream{ID:%d, X:%s, Y:%s, Cap:%d}", s.ID, s.X, s.Y, s.Cap)
}
