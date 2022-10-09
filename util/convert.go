package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func StringMustToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func StringMustToFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func StringMustToDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return d
}

func IntMustToDecimal(i int) decimal.Decimal {
	return decimal.NewFromInt32(int32(i))
}

func FloatMustToDecimal(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}
