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
	HardenIncrease  = 9.7
	plummetIncrease = -9.7
	errTime         = time.Time{}
)

func (r *repository) CreateStockPriceRanking(ranking model.Stock) error {
	return r.gormDB.Create(&ranking).Error
}

// 查询指定时间有涨停的票(去重复)
func (r *repository) ListStockPriceRanking(startTime, endTime time.Time) ([]model.Stock, error) {
	var spStocks []model.Stock
	err := r.gormDB.Group("stock_name").Where("increase_precent > ?", HardenIncrease).
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
	err := r.gormDB.Where("increase_precent > ? and create_time > ?", HardenIncrease, getTimeZero(lastTime)).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListTodayStockWithRose(start, end float64) ([]model.Stock, error) {
	var (
		spStocks []model.Stock
	)
	lastTime := r.GetLastCreateTime()
	err := r.gormDB.Where("increase_precent > ? and increase_precent < ? and create_time > ?", start, end, getTimeZero(lastTime)).Find(&spStocks).Error
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
	err := r.gormDB.Where("increase_precent > ? and create_time > ? and create_time < ?", HardenIncrease, startDay, endDay).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListPlummetStockLastDay() ([]model.Stock, error) {
	var (
		spStocks []model.Stock
	)
	lastTime := r.GetLastCreateTime()
	err := r.gormDB.Where("increase_precent < ? and create_time > ?", plummetIncrease, getTimeZero(lastTime)).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListTodayStockWithFall(start, end float64) ([]model.Stock, error) {
	var (
		spStocks []model.Stock
	)
	lastTime := r.GetLastCreateTime()
	err := r.gormDB.Where("increase_precent < ? and increase_precent > ? and create_time > ?", start, end, getTimeZero(lastTime)).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) ListPlummetStockWithDay(day int) ([]model.Stock, error) {
	var (
		lastStock model.Stock
		spStocks  []model.Stock
	)
	if err := r.gormDB.Last(&lastStock).Error; err != nil {
		return nil, err
	}
	endDay := getTimeZero(lastStock.CreateTime)
	startDay := getTimeEnd(lastStock.CreateTime.Add(-time.Duration(day*24) * time.Hour))
	err := r.gormDB.Where("increase_precent < ? and create_time > ? and create_time < ?", plummetIncrease, startDay, endDay).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetAllStock() ([]model.Stock, error) {
	last := r.GetLastCreateTime()
	var spStocks []model.Stock
	err := r.gormDB.Where("create_time = ?", last).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetAllSubordinate() ([]string, error) {
	time := r.GetLastCreateTime()
	var subordinates []string
	rows, err := r.gormDB.Raw(fmt.Sprintf("SELECT DISTINCT subordinate FROM %s WHERE create_time = '%s'",
		model.StockTableName,time)).Rows()
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

	lastTime := r.GetLastCreateTime()
	var spStocks []model.Stock
	err := r.gormDB.Where("create_time = ? and subordinate = ?", lastTime, subordinate).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}

func (r *repository) GetStockInfo(sc string) (*model.Stock, error) {

	lastTime := r.GetLastCreateTime()
	var spStocks model.Stock
	err := r.gormDB.Where("create_time = ? and stock_code = ?", lastTime, sc).Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return &spStocks, nil
}

func (r *repository) GetStockInfoLastDay(sc string, calTime string) ([]model.Stock, error) {

	var spStocks []model.Stock

	err := r.gormDB.Where("create_time > ? and create_time < ? and stock_code = ?", calTime, getTimeEnd(r.GetLastCreateTime()), sc).Order("id").Find(&spStocks).Error
	if err != nil {
		return nil, err
	}
	return spStocks, nil
}
