package server

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func EchoStock(list []model.Stock, less sortFun) string {
	if less != nil {
		CustomSortStock(list, less)
	}

	table, err := gotable.Create("行业", "代码", "名称", "市值", "pe", "涨幅","6日涨幅")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	for _, stock := range list {
		err := table.AddRow([]string{
			stock.Subordinate,
			stock.StockCode,
			stock.StockName,
			fmt.Sprintf("%.1f", stock.TotalMarketValue),
			fmt.Sprintf("%.1f", stock.Pe),
			fmt.Sprintf("%.1f", stock.IncreasePrecent),
			fmt.Sprintf("%.1f", stock.SixDaysUp),
		})
		if err != nil {
			log.Errorf("[table.AddRow] %v", err)
		}
	}
	return table.String()
}

type sortFun = func(x, y model.Stock) bool

var (
	SortWithSubordinatePe sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.Pe > y.Pe {
			return true
		}
		if x.Pe < y.Pe {
			return false
		}
		return true
	}
	SortWithSubordinateMarkValue sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.TotalMarketValue > y.TotalMarketValue {
			return true
		}
		if x.TotalMarketValue < y.TotalMarketValue {
			return false
		}
		return true
	}
	SortWithSubordinateSixDaysChange sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.SixDaysUp > y.SixDaysUp {
			return true
		}
		if x.SixDaysUp < y.SixDaysUp {
			return false
		}
		return true
	}

	SortWithSubordinateIncrease sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.IncreasePrecent > y.IncreasePrecent {
			return true
		}
		if x.IncreasePrecent < y.IncreasePrecent {
			return false
		}
		if x.TotalMarketValue > y.TotalMarketValue {
			return true
		}
		if x.TotalMarketValue < y.TotalMarketValue {
			return false
		}
		return true
	}

)

func CustomSortStock(source []model.Stock, less sortFun) {
	sort.Sort(customSort{
		source: source,
		less:   less,
	})
}

type customSort struct {
	source []model.Stock
	less   func(x, y model.Stock) bool
}

func (x customSort) Len() int           { return len(x.source) }
func (x customSort) Less(i, j int) bool { return x.less(x.source[i], x.source[j]) }
func (x customSort) Swap(i, j int)      { x.source[i], x.source[j] = x.source[j], x.source[i] }
