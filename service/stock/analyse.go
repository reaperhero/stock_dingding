package stock

import (
	"github.com/reaperhero/stock_dingding/model"
	"github.com/reaperhero/stock_dingding/model/repository"
	log "github.com/sirupsen/logrus"
)

func GetLastHardenStock() []model.Stock {
	list, err := repository.Repository.ListHardenStockLastDay()
	if err != nil {
		log.Errorf("[repository.Repository.ListHardenStockLastDay] %v", err)
	}
	return list
}

func GetHardenStockWithDays(day, count int) []model.Stock {
	list, err := repository.Repository.ListHardenStockWithDay(day)
	if err != nil {
		log.Errorf("[repository.Repository.ListHardenStockWithDay] %v", err)
	}
	m := make(map[string]int)
	for _, stock := range list {
		if v, ok := m[stock.StockCode]; ok {
			m[stock.StockCode] = v + 1
			continue
		}
		m[stock.StockCode] = 1
	}
	var result []model.Stock
	for mK, mV := range m {
		if mV >= count {
			for k := len(list) - 1; k >= 0; k-- {
				if list[k].StockCode == mK {
					result = append(result, list[k])
					break
				}
			}
		}
	}
	return result
}

func ChinaStockType() []model.Stock {
	list, err := repository.Repository.GetAllStock()
	if err != nil {
		log.Infoln(err)
		return nil
	}
	return list
}

func GetStockBySubordinate(subordinate string) ([]model.Stock, error) {
	list, err := repository.Repository.GetAllStockBySubordinate(subordinate)
	if err != nil {
		return nil, err
	}
	return list, nil
}
