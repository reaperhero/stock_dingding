package server

import (
	"fmt"
	"github.com/reaperhero/stock_dingding/service/stock"
	"github.com/reaperhero/stock_dingding/utils"
)

func reportDailyLimitStatisticsStock() {
	list := stock.GetLastHardenStock()
	fmt.Println(EchoStock(list, SortWithSubordinatePe))
}

func reportChinaAllStock() {
	list := stock.ChinaStockType()
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
