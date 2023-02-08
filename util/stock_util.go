package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

//单位转换
func UnitConversion(v string, u string) float64 {
	value := Substr(v, 0, len(v)-1)
	f, _ := strconv.ParseFloat(value, 64)
	d := decimal.NewFromFloat(f)
	unit := Substr(v, len(v)-1, len(v))
	var result float64 = f
	switch {
	case u == "万" && unit == "亿":
		result, _ = d.Div(decimal.NewFromFloat(10000)).Round(2).Float64()
	case u == "亿" && unit == "万":
		result, _ = d.Mul(decimal.NewFromFloat(10000)).Round(2).Float64()
	}
	return result
}
