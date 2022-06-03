package repository

import (
	"github.com/reaperhero/stock_dingding/model"
	"time"
)

var (
	hardenIncrease = 9.7
	errTime        = time.Time{}
)

func (r *repository) CreateStockPriceRanking(ranking model.Stock) error {
	return r.gormDB.Create(&ranking).Error
}

// 查询指定时间有涨停的票(去重复)
func (r *repository) ListStockPriceRanking(startTime, endTime time.Time) ([]model.Stock, error) {
	var spStocks []model.Stock
	err := r.gormDB.Group("stock_name").Where("increase_precent > ?", hardenIncrease).
		Where("create_time > ?", startTime).Where("create_time < ?", endTime).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListHardenStockLastDay() ([]model.Stock, error) {
	var (
		lastStock       model.Stock
		spStocks        []model.Stock
	)
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	err := r.gormDB.Where("increase_precent > ?", hardenIncrease).Where("from_days(to_days(create_time)) = ?", lastStock.CreateTime.Format("2006-01-02")).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetAllStock() ([]model.Stock, error) {

	var lastStock model.Stock
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	var spStocks []model.Stock
	err := r.gormDB.Where("from_days(to_days(create_time)) = ?", lastStock.CreateTime.Format("2006-01-02")).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetLastCreateTime() time.Time {

	var lastStock model.Stock
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return errTime
	}
	return lastStock.CreateTime
}

func (r *repository) GetAllStockBySubordinate(subordinate string) ([]model.Stock, error) {

	var lastStock model.Stock
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	var spStocks []model.Stock
	err := r.gormDB.Where("from_days(to_days(create_time)) = ?", lastStock.CreateTime.Format("2006-01-02")).Where("subordinate = ?", subordinate).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}
