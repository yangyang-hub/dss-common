package model

//同花顺概念列表
type LongHu struct {
	Symbol    string  `json:"symbol" gorm:"column:symbol;primary_key"`         //股票代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Name      string  `json:"name" gorm:"column:name"`                         //股票名称
	Close     float64 `json:"close" gorm:"column:close;float(11,8)"`           //收盘价
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(11,8)"`       //涨跌幅
	Amount    float64 `json:"amount" gorm:"column:amount;float(11,8)"`         //成交额(千元)
	Buy       float64 `json:"buy" gorm:"column:buy;float(11,8)"`               //净买入额(千元)
}

func (longHu LongHu) TableName() string {
	return "long_hu"
}

//同花顺概念列表
type ThsGn struct {
	Code string `json:"code" gorm:"column:code;primary_key"` //同花顺概念代码
	Name string `json:"name" gorm:"column:name"`             //同花顺概念名称
}

func (thsGn ThsGn) TableName() string {
	return "ths_gn"
}

//同花顺概念关联列表
type ThsGnRelSymbol struct {
	Symbol string `json:"symbol" gorm:"column:symbol"`   //股票代码
	GnName string `json:"gn_name" gorm:"column:gn_name"` //同花顺概念名称
}

func (thsGnRelSymbol ThsGnRelSymbol) TableName() string {
	return "ths_gn_rel_symbol"
}

//同花顺行业列表
type ThsHy struct {
	Code string `json:"code" gorm:"column:code;primary_key"` //同花顺行业代码
	Name string `json:"name" gorm:"column:name"`             //同花顺行业名称
}

func (thsHy ThsHy) TableName() string {
	return "ths_hy"
}

//同花顺行业关联列表
type ThsHyRelSymbol struct {
	HyCode string `json:"hy_code" gorm:"column:hy_code;primary_key"` //同花顺行业代码
	Symbol string `json:"symbol" gorm:"column:symbol;primary_key"`   //股票代码
}

func (thsHyRelSymbol ThsHyRelSymbol) TableName() string {
	return "ths_hy_rel_symbol"
}

//同花顺行业行情
type ThsHyQuote struct {
	Code      string  `json:"code" gorm:"column:code;primary_key"`             //行业代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(30,4)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(30,4)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(30,4)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(30,4)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(30,4)"`   //昨收价
	Change    float64 `json:"change" gorm:"column:change;float(30,4)"`         //资金流入(亿)
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(30,4)"`       //涨跌幅
	Vol       float64 `json:"vol" gorm:"column:vol;float(30,4)"`               //成交量(万手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(30,4)"`         //成交额(亿)
	Rank      int     `json:"rank" gorm:"column:rank;int(11)"`                 //涨幅排名
	RiseCount int     `json:"rise_count" gorm:"column:rise_count;int(11)"`     //上涨家数
	FallCount int     `json:"fall_count" gorm:"column:fall_count;int(11)"`     //下跌家数
}

func (thsHyQuote ThsHyQuote) TableName() string {
	return "ths_hy_quote"
}

//同花顺概念行情
type ThsGnQuote struct {
	Code      string  `json:"code" gorm:"column:code;primary_key"`             //概念代码
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(30,4)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(30,4)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(30,4)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(30,4)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(30,4)"`   //昨收价
	Change    float64 `json:"change" gorm:"column:change;float(30,4)"`         //资金流入(亿)
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(30,4)"`       //涨跌幅
	Vol       float64 `json:"vol" gorm:"column:vol;float(30,4)"`               //成交量(万手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(30,4)"`         //成交额(亿)
	Rank      int     `json:"rank" gorm:"column:rank;int(11)"`                 //涨幅排名
	RiseCount int     `json:"rise_count" gorm:"column:rise_count;int(11)"`     //上涨家数
	FallCount int     `json:"fall_count" gorm:"column:fall_count;int(11)"`     //下跌家数
}

func (thsGnQuote ThsGnQuote) TableName() string {
	return "ths_gn_quote"
}
