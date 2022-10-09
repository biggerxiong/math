package model

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	pa := &Point{
		X: 23.127613278533207,
		Y: 113.36619463562009,
	}
	pb := &Point{
		X: 23.125876748465735,
		Y: 113.38340368865964,
	}

	fmt.Println(pa.Distance(pb))
}
