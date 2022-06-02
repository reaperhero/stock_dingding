package server

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/reaperhero/stock_dingding/model"
	"github.com/reaperhero/stock_dingding/service/stock_analyse"
	"github.com/reaperhero/stock_dingding/utils"
	log "github.com/sirupsen/logrus"
)

func reportDailyLimitStatisticsStock() {
	list := stock_analyse.GetLastHardenStock()
	sortStock(list)

	echoStock(list)
}

func reportChinaAllStock() {
	list := stock_analyse.ChinaStockType()
	var m = make(map[string]*int)
	for _, ranking := range list {
		if _, ok := m[ranking.Subordinate]; ok {
			*m[ranking.Subordinate]++
			continue
		}
		var in = new(int)
		m[ranking.Subordinate] = in
	}

	var (
		source = make(map[interface{}]interface{})
	)
	for s, i := range m {
		source[s] = *i
	}
	ks := utils.SortMapWithValue(source, true)
	for _, k := range ks {
		fmt.Printf("[%v]: %v只股\n", k, source[k])
	}
}

func echoStock(list []model.Stock) {
	//fmt.Printf("%-30s %-30s %-30s %-30s %-30s %-30s \n", "行业", "代码", "名称", "市值", "pe", "涨幅")
	//
	//for _, stock := range list {
	//	fmt.Printf("%-30s %-30s %-30s %-30.1f %-30.1f %-30.1f \n", stock.Subordinate, stock.StockCode, stock.StockName, stock.TotalMarketValue, stock.Pe, stock.IncreasePrecent)
	//}

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
