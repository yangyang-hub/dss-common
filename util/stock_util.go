package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// 计算涨停板
func CalLimitUp(code string, preClose, change float64, stStock map[string]string) int64 {
	prefix := Substr(code, 0, 2)
	var divValue float64 = 10
	if stStock[code] != "" {
		divValue = 5
	} else if prefix == "60" {
		divValue = 10
	} else if prefix == "00" {
		divValue = 10
	} else if prefix == "68" {
		divValue = 20
	} else if prefix == "30" {
		divValue = 20
	} else if Substr(code, 0, 1) == "8" {
		divValue = 30
	}
	preClosed := decimal.NewFromFloat(preClose)
	p := decimal.NewFromFloat(divValue)
	limit := preClosed.Div(p).Round(2)
	res := decimal.NewFromFloat(change).Cmp(limit)
	if res >= 0 {
		return 1
	} else {
		return 0
	}
}

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
