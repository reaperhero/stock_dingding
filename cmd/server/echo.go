package server

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

func echoStock(list []model.Stock, less sortFun) {
	if less != nil {
		customSortStock(list, less)
	}

	table, err := gotable.Create("行业", "代码", "名称", "市值", "pe", "涨幅")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, stock := range list {
		err := table.AddRow([]string{
			stock.Subordinate,
			stock.StockCode,
			stock.StockName,
			fmt.Sprintf("%.1f", stock.TotalMarketValue),
			fmt.Sprintf("%.1f", stock.Pe),
			fmt.Sprintf("%.1f", stock.IncreasePrecent),
		})
		if err != nil {
			log.Errorf("[table.AddRow] %v", err)
		}
	}
	fmt.Println(table)
}

type sortFun = func(x, y model.Stock) bool

var (
	sortWithSubordinatePe sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.Pe >= y.Pe {
			return false
		}
		return true
	}
	sortWithSubordinateMarkValue sortFun = func(x, y model.Stock) bool {
		if strings.Compare(x.Subordinate, y.Subordinate) > 0 {
			return true
		}
		if strings.Compare(x.Subordinate, y.Subordinate) < 0 {
			return false
		}
		if x.TotalMarketValue >= y.TotalMarketValue {
			return false
		}
		return true
	}
)

func customSortStock(source []model.Stock, less sortFun) {
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
