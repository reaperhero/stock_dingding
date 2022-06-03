package stock

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/model"
	"github.com/reaperhero/stock_dingding/model/repository"
	log "github.com/sirupsen/logrus"
	"sort"
)

// 最近7天某天涨幅超过9%，统计次数
// 农牧饲渔 [罗 牛 山+涨停1次 敦煌种业+涨停1次]
func DailyLimitStatistics(searchDay string) (resultStock map[string][]string, maxCount int) {
	list, _ := repository.Repository.ListHardenStockLastDay()

	var (
		m = make(map[string]int)
	)
	for _, ranking := range list {
		m[ranking.StockCode]++
	}
	type rank struct {
		StockCode       string
		StockName       string
		IncreasePrecent float64
		IncreaseSpeed   float64
		Count           int
		Subordinate     string
	}

	var result []rank // 股票+涨停次数
	for k, count := range m {
		for _, ranking := range list {
			if ranking.StockCode == k && ranking.CreateTime.Format("2006-01-02") == searchDay {
				result = append(result, rank{
					StockCode:       ranking.StockCode,
					StockName:       ranking.StockName,
					IncreasePrecent: ranking.IncreasePrecent,
					IncreaseSpeed:   ranking.IncreaseSpeed,
					Count:           count,
					Subordinate:     ranking.Subordinate,
				})
				break
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})
	hangMap := make(map[string][]string) // 行业  股票池
	for _, r := range result {
		code := fmt.Sprintf("%s+涨停%d次", r.StockName, r.Count)
		hangMap[r.Subordinate] = append(hangMap[r.Subordinate], code)
	}
	for _, stocks := range hangMap {
		if len(stocks) > maxCount {
			maxCount = len(stocks)
		}
	}

	return hangMap, maxCount
}

func GetLastHardenStock() []model.Stock {
	list, err := repository.Repository.ListHardenStockLastDay()
	if err != nil {
		log.Errorf("[repository.Repository.ListHardenStockLastDay] %v", err)
	}
	return list
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
