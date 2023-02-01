package model

//实时数据
type LiveData struct {
	Code          string  `json:"code"`          //编码
	Name          string  `json:"name"`          //名称
	Now           float64 `json:"now"`           //当前价
	Change        float64 `json:"change"`        //涨跌
	ChangePercent float64 `json:"changePercent"` //涨跌幅
	Time          string  `json:"time"`          //更新时间
	Max           float64 `json:"max"`           //最高价
	Min           float64 `json:"min"`           //最低价
}