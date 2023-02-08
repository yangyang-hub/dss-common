package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

//单位转换为万
func UnitConversion(v string) float64 {
	value := Substr(v, 0, len(v)-1)
	f, _ := strconv.ParseFloat(value, 64)
	d := decimal.NewFromFloat(f)
	unit := Substr(v, len(v)-1, len(v))
	var result float64 = f
	switch {
	case unit == "亿":
		result, _ = d.Div(decimal.NewFromFloat(10000)).Round(2).Float64()
	case unit == "万":
		result = f
	}
	return result
}
