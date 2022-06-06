package server

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sort"
)

func compareFun(src, dst sortFun) bool {
	sf1 := reflect.ValueOf(src)
	sf2 := reflect.ValueOf(dst)
	return sf1.Pointer() == sf2.Pointer()
}
func EchoStock(list []model.Stock, less sortFun) string {

	if less != nil {
		CustomSortStock(list, less)
	}

	table, err := gotable.Create("行业", "代码", "名称", "市值", "pe", "涨幅", "3日涨幅", "6日涨幅")
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
			fmt.Sprintf("%.1f", stock.ThreeDaysUp),
			fmt.Sprintf("%.1f", stock.SixDaysUp),
		})
		if err != nil {
			log.Errorf("[table.AddRow] %v", err)
		}
	}
	result := table.String()
	switch {
	case compareFun(less, SortWithSubordinateMarkValue):
		result = fmt.Sprintln("[行业>市值]") + result
	case compareFun(less, SortWithSubordinatePe):
		result = fmt.Sprintln("[行业>PE]") + result
	case compareFun(less, SortWithSubordinateThreeDaysChange):
		result = fmt.Sprintln("[行业>三日涨幅]") + result
	case compareFun(less, SortWithSubordinateIncrease):
		result = fmt.Sprintln("[行业>单日涨幅]") + result
	case compareFun(less, SortWithSubordinateSixDaysChange):
		result = fmt.Sprintln("[行业>六日涨幅]") + result
	}
	return result
}

type sortFun = func(x, y model.Stock) bool

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
