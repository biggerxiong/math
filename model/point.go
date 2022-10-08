package model

import (
	"main/util"
	"math"
)

type Point struct {
	X string
	Y string
}

// 地球的平均半径以千米为单位
const radius = 6371

// degrees2radians 度数转弧度公式
func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Distance Haversine公式(半正矢公式):计算两点之间的直线距离
func (origin *Point) Distance(destination *Point) float64 {
	fOriginX := util.StringMustToFloat(origin.X)
	fOriginY := util.StringMustToFloat(origin.Y)
	fDestX := util.StringMustToFloat(destination.X)
	fDestY := util.StringMustToFloat(destination.Y)

	degreesLat := degrees2radians(fDestX - fOriginX)
	degreesLong := degrees2radians(fDestY - fOriginY)
	a := math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(fOriginX))*
			math.Cos(degrees2radians(fDestX))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return d
}
