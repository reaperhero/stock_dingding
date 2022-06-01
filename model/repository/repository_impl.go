package repository

import (
	"github.com/reaperhero/stock_dingding/model"
	"time"
)

func (r *repository) CreateStockPriceRanking(ranking model.StockPriceRanking) error {
	return r.gormDB.Create(&ranking).Error
}

func (r *repository) TodayStockRanking() []model.StockPriceRanking {
	var (
		spStocks  []model.StockPriceRanking
		startTime = time.Now()
	)
	if err := r.gormDB.Where("from_days(to_days(create_time)) = ?", startTime.Format("2006-01-02")).Find(&spStocks).Error; err != nil {
		return nil
	}
	return spStocks
}

func (r *repository) ListStockPriceRanking(startTime, endTime time.Time, increasePrecent float64) ([]model.StockPriceRanking, error) {
	var spStocks []model.StockPriceRanking
	err := r.gormDB.Where("increase_precent > ?", increasePrecent).Where("create_time > ?", startTime).Where("create_time < ?", endTime).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetAllStock() ([]model.StockPriceRanking, error) {

	var lastStock model.StockPriceRanking
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	var spStocks []model.StockPriceRanking
	err := r.gormDB.Where("from_days(to_days(create_time)) = ?", lastStock.CreateTime.Format("2006-01-02")).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}
