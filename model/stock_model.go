package model

import (
	"fmt"

	util "github.com/yangyang-hub/dss-common/util"
)

type Stock interface {
	StockInfo | StockCompany | StockQuote
}

type Tabler interface {
	TableName() string
}

//交易日历
type TradeCal struct {
	Exchange     string `json:"exchange"`      //交易所 SSE上交所 SZSE深交所
	CalDate      string `json:"cal_date"`      //日历日期
	IsOpen       int64  `json:"is_open"`       //是否交易 0休市 1交易
	PretradeDate string `json:"pretrade_date"` //上一个交易日
}

//基础数据
type StockInfo struct {
	TsCode     string `json:"ts_code" gorm:"column:ts_code;primary_key"` //TS代码
	Symbol     string `json:"symbol" gorm:"column:symbol"`               //股票代码
	Name       string `json:"name" gorm:"column:name"`                   //股票名称
	Area       string `json:"area" gorm:"column:area"`                   //地域
	Industry   string `json:"industry" gorm:"column:industry"`           //所属行业
	Fullname   string `json:"fullname" gorm:"column:fullname"`           //股票全称
	Enname     string `json:"enname" gorm:"column:enname"`               //英文全称
	Cnspell    string `json:"cnspell" gorm:"column:cnspell"`             //拼音缩写
	Market     string `json:"market" gorm:"column:market"`               //市场类型（主板/创业板/科创板/CDR）
	Exchange   string `json:"exchange" gorm:"column:exchange"`           //交易所代码
	CurrType   string `json:"curr_type" gorm:"column:curr_type"`         //交易货币
	ListStatus string `json:"list_status" gorm:"column:list_status"`     //上市状态 L上市 D退市 P暂停上市
	ListDate   string `json:"list_date" gorm:"column:list_date"`         //上市日期
	DelistDate string `json:"delist_date" gorm:"column:delist_date"`     //退市日期
	IsHs       string `json:"is_hs" gorm:"column:is_hs"`                 //是否沪深港通标的，N否 H沪股通 S深股通
}

func (stockInfo StockInfo) TableName() string {
	//表名
	return "stock_info"
}

//上市公司基本信息
type StockCompany struct {
	TsCode        string  `json:"ts_code" gorm:"column:ts_code;primary_key"`           //TS代码
	Exchange      string  `json:"exchange" gorm:"column:exchange"`                     //交易所代码
	Chairman      string  `json:"chairman" gorm:"column:chairman"`                     //法人代表
	Manager       string  `json:"manager" gorm:"column:manager"`                       //总经理
	Secretary     string  `json:"secretary" gorm:"column:secretary"`                   //董秘
	RegCapital    float64 `json:"reg_capital" gorm:"column:reg_capital;decimal(11,8)"` //注册资本
	SetupDate     string  `json:"setup_date" gorm:"column:setup_date"`                 //注册日期
	Province      string  `json:"province" gorm:"column:province"`                     //所在省份
	City          string  `json:"city" gorm:"column:city"`                             //所在城市
	Introduction  string  `json:"introduction" gorm:"column:introduction"`             //公司介绍
	Website       string  `json:"website" gorm:"column:website"`                       //公司主页
	Email         string  `json:"email" gorm:"column:email"`                           //电子邮件
	Office        string  `json:"office" gorm:"column:office"`                         //办公室
	Employees     int64   `json:"employees" gorm:"column:employees"`                   //员工人数
	MainBusiness  string  `json:"main_business" gorm:"column:main_business"`           //主要业务及产品
	BusinessScope string  `json:"business_scope" gorm:"column:business_scope"`         //经营范围
}

func (stockCompany StockCompany) TableName() string {
	//表名
	return "stock_company"
}

//股票行情
type StockQuote struct {
	TsCode    string  `json:"ts_code" gorm:"column:ts_code;primary_key"`       //股票代码
	Type      string  `json:"type" gorm:"column:type;primary_key"`             //D：日线行情；W：周线行情；M：月线行情
	TradeDate string  `json:"trade_date" gorm:"column:trade_date;primary_key"` //交易日期
	Open      float64 `json:"open" gorm:"column:open;float(11,8)"`             //开盘价
	High      float64 `json:"high" gorm:"column:high;float(11,8)"`             //最高价
	Low       float64 `json:"low" gorm:"column:low;float(11,8)"`               //最低价
	Close     float64 `json:"close" gorm:"column:close;float(11,8)"`           //收盘价
	PreClose  float64 `json:"pre_close" gorm:"column:pre_close;float(11,8)"`   //昨收价(前复权)
	Change    float64 `json:"change" gorm:"column:change;float(11,8)"`         //涨跌额
	PctChg    float64 `json:"pct_chg" gorm:"column:pct_chg;float(11,8)"`       //涨跌幅(未复权)
	Vol       float64 `json:"vol" gorm:"column:vol;float(11,8)"`               //成交量(手)
	Amount    float64 `json:"amount" gorm:"column:amount;float(11,8)"`         //成交额(千元)
}

func (stockQuote StockQuote) TableName() string {
	//表名
	tableName := fmt.Sprintf("%s%s", "stock_quote_", util.Substr(stockQuote.TradeDate, 0, 4))
	return tableName
}
