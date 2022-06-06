package repository

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func getTimeZero(t time.Time) string {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startTime.Format("2006/01/02 15:04:05")
}

func getTimeEnd(t time.Time) string {
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return endTime.Format("2006/01/02 15:04:05")

}

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
		spStocks []model.Stock
	)
	lastTime := r.GetLastCreateTime()
	err := r.gormDB.Where("increase_precent > ? and create_time > ?", hardenIncrease, getTimeZero(lastTime)).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListHardenStockWithDay(day int) ([]model.Stock, error) {
	var (
		lastStock model.Stock
		spStocks  []model.Stock
	)
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	endDay := getTimeZero(lastStock.CreateTime)
	startDay := getTimeEnd(lastStock.CreateTime.Add(-time.Duration(day*24) * time.Hour))
	err := r.gormDB.Where("increase_precent > ? and create_time > ? and create_time < ?", hardenIncrease, startDay, endDay).Find(&spStocks).Error
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

func (r *repository) GetAllSubordinate() ([]string, error) {
	var lastStock model.Stock
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	var subordinates []string
	rows, err := r.gormDB.Raw(fmt.Sprintf("SELECT DISTINCT subordinate FROM %s WHERE from_days(to_days(create_time)) = '%s'",
		model.StockTableName, lastStock.CreateTime.Format("2006-01-02"))).Rows()
	if err != nil {
		return nil, err
	}
	//rows, err := r.gormDB.Model(&model.Stock{}).Where("from_days(to_days(create_time)) = ?",lastStock.CreateTime.Format("2006-01-02")).Select("subordinate").Rows()

	defer rows.Close()
	if err != nil {
		log.Errorf("[repository.GetAllSubordinate] %v", err)
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Errorf("[repository.GetAllSubordinate] rows.Scan %v", err)
			continue
		}
		subordinates = append(subordinates, name)
	}
	return subordinates, nil
}

func (r *repository) GetRankSubordinate(subordinate string) ([]model.Stock, error) {
	lastTime := r.GetLastCreateTime()
	var stocks []model.Stock
	err := r.gormDB.Where("from_days(to_days(create_time))= ? and subordinate = ?", lastTime.Format("2006-01-02"), subordinate).Find(&stocks).Error
	if err != nil {
		return nil, err
	}
	return stocks, nil
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
