package model

import "fmt"

type Node struct {
	ID int
	X  string
	Y  string
}

func (p Node) String() string {
	return fmt.Sprintf("Node{X: %s, Y:%s}", p.X, p.Y)
}

func (p Node) Key() string {
	return p.String()
}
