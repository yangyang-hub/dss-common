package constant

//批量新增每批次数量
const InsertBatchSize int = 1000

//交易所编码枚举
type Exchange int

const ExchangeConst Exchange = 0

func (n Exchange) String() string {
	return [...]string{"SSE", "SZSE", "BSE"}[n]
}

func (n Exchange) List() []string {
	return []string{"SSE", "SZSE", "BSE"}
}

func (n Exchange) Exclude(i Exchange) []string {
	exchange := ExchangeConst.List()
	exchange = append(exchange[:i], exchange[i+1:]...)
	return exchange
}

const (
	SSE  Exchange = iota //上交所
	SZSE                 //深交所
	BSE                  //北交所
)

//行情信息枚举
type Quote int

const QuoteConst Quote = 0

func (n Quote) String() string {
	return [...]string{"daily", "weekly", "monthly"}[n]
}

func (n Quote) List() []string {
	return []string{"daily", "weekly", "monthly"}
}

func (n Quote) Exclude(i Quote) []string {
	quote := QuoteConst.List()
	quote = append(quote[:i], quote[i+1:]...)
	return quote
}

const (
	daily   Quote = iota //日行情
	weekly               //周行情
	monthly              //月行情
)
