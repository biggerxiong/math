package intf

import "main/model"

type Key interface {
	Key() string
}

type Pointer interface {
	GetPoint() *model.Point
}
