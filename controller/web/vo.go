package web

type mStcok struct {
	StockCode           string    `db:"stock_code" json:"stock_code"`                       //  代码
	StockName           string    `db:"stock_name" json:"stock_name"`                       //  名称
	CurentPrice         float64   `db:"curent_price" json:"curent_price"`                   //  最新
	Pe                  float64   `db:"pe" json:"pe"`                                       //  市盈率
	Subordinate         string    `db:"subordinate" json:"subordinate"`                     //  所属行业
	TotalMarketValue    float64   `db:"total_market_value" json:"total_market_value"`       //  总市值 亿
	ThreeDaysUp         float64   `db:"three_days_up" json:"three_days_up"`                 //  3日涨幅
	SixDaysUp           float64   `db:"six_days_up" json:"six_days_up"`                     //  6日涨幅
	ThreeDaysChange     float64   `db:"three_days_change" json:"three_days_change"`         //  3日换手
}
