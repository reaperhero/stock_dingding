package model

import "time"

type StockPriceRanking struct {
	ID                  int64     `db:"id" json:"id"`
	CreateTime          time.Time `db:"create_time" json:"create_time"`
	StockCode           string    `db:"stock_code" json:"stock_code"`                       //  代码
	StockName           string    `db:"stock_name" json:"stock_name"`                       //  名称
	CurentPrice         float64   `db:"curent_price" json:"curent_price"`                   //  最新
	IncreasePrecent     float64   `db:"increase_precent" json:"increase_precent"`           //  涨幅
	QuantumTotal        float64   `db:"quantum_total" json:"quantum_total"`                 //  总量 万
	QuantumCurrent      float64   `db:"quantum_current" json:"quantum_current"`             //  现量 万
	BuyPrice            float64   `db:"buy_price" json:"buy_price"`                         //  买入价
	SellPrice           float64   `db:"sell_price" json:"sell_price"`                       //  卖出价
	IncreaseSpeed       float64   `db:"increase_speed" json:"increase_speed"`               //  涨速
	Turnover            float64   `db:"turnover" json:"turnover"`                           //  换手
	Amount              float64   `db:"amount" json:"amount"`                               //  金额 亿
	Pe                  float64   `db:"pe" json:"pe"`                                       //  市盈率
	Subordinate         string    `db:"subordinate" json:"subordinate"`                     //  所属行业
	Highest             float64   `db:"highest" json:"highest"`                             //  最高
	Minimum             float64   `db:"minimum" json:"minimum"`                             //  最低
	OpenPrice           float64   `db:"open_price" json:"open_price"`                       //  开盘
	YesterdayClosePrice float64   `db:"yesterday_close_price" json:"yesterday_close_price"` //  昨收
	Amplitude           float64   `db:"amplitude" json:"amplitude"`                         //  振幅
	Appoint             float64   `db:"appoint" json:"appoint"`                             //  量比
	Magnitude           float64   `db:"magnitude" json:"magnitude"`                         //  委比
	Committee           float64   `db:"committee" json:"committee"`                         //  委差 万
	Average             float64   `db:"average" json:"average"`                             //  均价
	InsideDish          float64   `db:"inside_dish" json:"inside_dish"`                     //  内盘 万
	OutsideDish         float64   `db:"outside_dish" json:"outside_dish"`                   //  外盘 万
	SideThan            float64   `db:"side_than" json:"side_than"`                         //  内外比
	BuyQuantity         float64   `db:"buy_quantity" json:"buy_quantity"`                   //  买一量 万
	SellQuantity        float64   `db:"sell_quantity" json:"sell_quantity"`                 //  卖一量 万
	PriceToBook         float64   `db:"price_to_book" json:"price_to_book"`                 //  市净率
	TotalEquity         float64   `db:"total_equity" json:"total_equity"`                   //  总股本 亿
	TotalMarketValue    float64   `db:"total_market_value" json:"total_market_value"`       //  总市值 亿
	CirculationCapital  float64   `db:"circulation_capital" json:"circulation_capital"`     //  流通股本 亿
	CurrentMarket       float64   `db:"current_market" json:"current_market"`               //  流通市值 亿
	ThreeDaysUp         float64   `db:"three_days_up" json:"three_days_up"`                 //  3日涨幅
	SixDaysUp           float64   `db:"six_days_up" json:"six_days_up"`                     //  6日涨幅
	ThreeDaysChange     float64   `db:"three_days_change" json:"three_days_change"`         //  3日换手
	SixDaysChange       float64   `db:"six_days_change" json:"six_days_change"`             //  6日换手
	EvenNumberDays      float64   `db:"even_number_days" json:"even_number_days"`           //  连涨天数
	UpThisMonth         float64   `db:"up_this_month" json:"up_this_month"`                 //  本月涨幅
	UpThisYear          float64   `db:"up_this_year" json:"up_this_year"`                   //  今年涨幅
	UpLastMonth         float64   `db:"up_last_month" json:"up_last_month"`                 //  近一月涨幅
	UpPastYear          float64   `db:"up_past_year" json:"up_past_year"`                   //  近一年涨幅
}
