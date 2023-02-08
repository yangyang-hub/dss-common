package util

import (
	"errors"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

func FloatCompare(f1, f2 interface{}) (n int, err error) {
	var f1Dec, f2Dec decimal.Decimal
	switch f1.(type) {
	case float64:
		f1Dec = decimal.NewFromFloat(f1.(float64))
		switch f2.(type) {
		case float64:
			f2Dec = decimal.NewFromFloat(f2.(float64))
		case string:
			f2Dec, err = decimal.NewFromString(f2.(string))
			if err != nil {
				return 2, err
			}
		default:
			return 2, errors.New("FloatCompare() expecting to receive float64 or string")
		}
	case string:
		f1Dec, err = decimal.NewFromString(f1.(string))
		if err != nil {
			return 2, err
		}
		switch f2.(type) {
		case float64:
			f2Dec = decimal.NewFromFloat(f2.(float64))
		case string:
			f2Dec, err = decimal.NewFromString(f2.(string))
			if err != nil {
				return 2, err
			}
		default:
			return 2, errors.New("FloatCompare() expecting to receive float64 or string")
		}
	default:
		return 2, errors.New("FloatCompare() expecting to receive float64 or string")
	}
	return f1Dec.Cmp(f2Dec), nil
}

// 强制舍弃尾数
func FormatFloatFloor(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Floor(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)

}

// 舍弃的尾数不为0，强制进位
func FormatFloatCeil(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Ceil(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)

}

// 保留两位小数，舍弃尾数，无进位运算
// 主要逻辑就是先乘，trunc之后再除回去，就达到了保留N位小数的效果
func FormatFloat(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)

}
