package model

import "fmt"

type Node struct {
	ID int
	Point
}

func (p Node) String() string {
	return fmt.Sprintf("Node{X: %s, Y:%s}", p.X, p.Y)
}

func (p Node) Key() string {
	return p.String()
}

func (p Node) GetPoint() *Point {
	return &p.Point
}
