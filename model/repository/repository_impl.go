package repository

import (
	"github.com/reaperhero/stock_dingding/model"
	"time"
)

func (r *repository) CreateStockPriceRanking(ranking model.StockPriceRanking) error {
	return r.gormDB.Create(&ranking).Error
}

func (r *repository) ListStockPriceRanking(startTime,endTime time.Time,increasePrecent float64) ([]model.StockPriceRanking, error) {
	var spStocks []model.StockPriceRanking
	err := r.gormDB.Where("increase_precent > ?", increasePrecent).Where("create_time > ?",startTime).Where("create_time < ?",endTime).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}
