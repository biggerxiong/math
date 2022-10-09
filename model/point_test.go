package model

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	pa := &Point{
		X: 1,
		Y: 2,
	}
	pb := &Point{
		X: 2,
		Y: 6,
	}

	fmt.Println(pa.Distance(pb))
}
