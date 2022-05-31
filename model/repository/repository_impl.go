package repository

import "github.com/reaperhero/stock_dingding/model"

func (r *repository) CreateStockPriceRanking(ranking model.StockPriceRanking) error {
	return r.gormDB.Create(&ranking).Error
}
