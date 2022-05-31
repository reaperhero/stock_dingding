package stock_analyse

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/model/repository"
	"sort"
	"time"
)

func DailyLimitStatistics() (resultStock map[string][]string,maxCount int){
	list, _ := repository.Repository.ListStockPriceRanking(time.Now().Add(-time.Hour*24*7), time.Now(), 9)
	sort.Slice(list, func(i, j int) bool {
		return list[i].IncreaseSpeed > list[j].IncreaseSpeed
	})
	m := make(map[string]int)

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
			if ranking.StockCode == k && ranking.CreateTime.Day() == 31 {
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

	return hangMap,maxCount
}
