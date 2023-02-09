package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

//单位转换为万
func UnitConversion(v string) float64 {
	l := len([]rune(v))
	value := Substr(v, 0, (l - 2))
	f, _ := strconv.ParseFloat(value, 64)
	d := decimal.NewFromFloat(f)
	unit := Substr(v, l-1, l)
	var result float64 = f
	if unit == "亿" {
		result, _ = d.Mul(decimal.NewFromFloat(10000)).Round(2).Float64()
	}
	return result
}
