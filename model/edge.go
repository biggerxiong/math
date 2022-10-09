package model

import "fmt"

type Edge struct {
	ID   int
	From int
	To   int
	Dis  float64
}

func (e Edge) String() string {
	return fmt.Sprintf("Edge{ID:%d, From:%d, To:%d, Dis:%f}", e.ID, e.From, e.To, e.Dis)
}
