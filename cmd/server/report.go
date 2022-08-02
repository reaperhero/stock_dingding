package server

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/reaperhero/stock_dingding/model"
	"github.com/reaperhero/stock_dingding/model/repository"
	"github.com/reaperhero/stock_dingding/service/stock"
	"github.com/reaperhero/stock_dingding/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

func reportCareAboutStock(day, count int) {
	list := stock.GetHardenStockWithDays(day, count)
	fmt.Println(EchoStock(list, SortWithSubordinateThreeDaysChange))
	fmt.Println(EchoStock(list, SortWithSubordinateMarkValue))
}

func reportDailyLimitStatisticsStock() {
	list := stock.GetLastHardenStock()
	fmt.Println(EchoStock(list, SortWithSubordinateMarkValue))
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

func reportCareAboutStockTofile() {

	dir := "docs/" + time.Now().Format("20060102") + "/"
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatalf("[reportCareAboutStockTofile] os.RemoveAll %v", err)
	}
	err = os.MkdirAll(dir+"subordinate", os.ModePerm)
	if err != nil {
		log.Fatalf("[reportCareAboutStockTofile] os.MkdirAll %v", err)
	}

	//list := stock.GetHardenStockWithDays(7, 2)
	//content := EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内2次涨停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetHardenStockWithDays(7, 3)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内3次涨停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetHardenStockWithDays(7, 4)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内4次涨停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetLastRoseStock(3, 6)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"今日涨幅在3-6.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetLastHardenStock()
	//hadenStockMap := make(map[string][]model.Stock)
	//for _, hadenStock := range list {
	//	if _, ok := hadenStockMap[hadenStock.Subordinate]; ok {
	//		hadenStockMap[hadenStock.Subordinate] = append(hadenStockMap[hadenStock.Subordinate], hadenStock)
	//		continue
	//	}
	//	hadenStockMap[hadenStock.Subordinate] = []model.Stock{hadenStock}
	//}
	//
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"今日涨停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetPlummetStockWithDays(7, 2)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内2次跌停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetPlummetStockWithDays(7, 3)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内3次跌停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetPlummetStockWithDays(7, 4)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"7日内4次跌停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetLastFalltStock(-3, -6)
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"今日跌幅在3-6.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//list = stock.GetLastPlummetStock()
	//plummetStockMap := make(map[string][]model.Stock)
	//for _, plummetStock := range list {
	//	if _, ok := hadenStockMap[plummetStock.Subordinate]; ok {
	//		plummetStockMap[plummetStock.Subordinate] = append(plummetStockMap[plummetStock.Subordinate], plummetStock)
	//		continue
	//	}
	//	plummetStockMap[plummetStock.Subordinate] = []model.Stock{plummetStock}
	//}
	//content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	//if err := ioutil.WriteFile(dir+"今日跌停.txt", []byte(content), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}
	//
	//if err := ioutil.WriteFile(dir+"龙头定位.txt", []byte(formatMap(hadenStockMap, plummetStockMap)), 0644); err != nil {
	//	log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	//}

	ts, err := repository.Repository.GetAllSubordinate()
	if err != nil {
		log.Fatal("[repository.Repository.GetAllSubordinate] %v", err)
	}

	for _, t := range ts {
		stocks, err := repository.Repository.GetAllStockBySubordinate(t)
		if err != nil {
			log.Fatalf("[repository.Repository.GetAllSubordinate] %v", err)
		}

		data := fmt.Sprintf("%s \n\n\n\n\n %s",
			EchoStock(stocks, SortWithSubordinateMarkValue),
			EchoStock(stocks, SortWithSubordinateIncrease),
		)

		err = ioutil.WriteFile(dir+"/subordinate/"+t+".txt", []byte(data), 0644)
		if err != nil {
			log.Fatalf("ioutil.WriteFile %v", err)
		}
	}
}

var (
	stionDayRecord     = 30 // record 保留的时间
	stionDayRecordInit = 7  // record 出事记录7天类有涨停的股票
	stionDayReport     = 10 // 报告统计天数
)

func trendStock() {

	m := stock.NewSitonManage()
	if !m.IsTodayRecord() {
		source := stock.GetLastHardenStock()

		soMap := hashmap.New()
		for _, m := range source {
			soMap.Put(m.StockCode, m)
		}
		// 查询record文件中，今天的涨幅
		for _, s := range m.GetStockSet() {
			info, err := repository.Repository.GetStockInfo(s)
			if err != nil {
				continue
			}
			soMap.Put(info.StockCode, *info)
		}
		relSource := make([]model.Stock, 0, 100)
		for _, v := range soMap.Values() {
			relSource = append(relSource, v.(model.Stock))
		}
		m.AddTodayStock(relSource) // 内存中整合今天的今天涨停的股票
	}
	m.RecordFile(stionDayRecord) // 数据保留多少天
	m.ReportFile(stionDayReport)
}

func initTrendStock() {

	var (
		res []stock.Siton
	)
	sts := stock.GetHardenStockWithDays(stionDayRecordInit, 1)

	for _, st := range sts {
		d := time.Now().Add(-time.Duration(stionDayRecord*24) * time.Hour).Format("2006-01-02")
		codeStocks, err := repository.Repository.GetStockInfoLastDay(st.StockCode, d)
		if err != nil {
			log.Errorf("[initTrendStock] repository.Repository.GetStockInfoLastDay %v", err)
		}
		ton := stock.Siton{
			StockCode:   st.StockCode,
			StockName:   st.StockName,
			Increases:   []float64{},
			Subordinate: st.Subordinate,
			CreateTime:  stock.StionTime,
		}
		for _, codeStock := range codeStocks {
			ton.Increases = append(ton.Increases, codeStock.IncreasePrecent)
		}

		res = append(res, ton)
	}

	stock.RecordFileWithStions(res)

	m := stock.NewSitonManage()

	m.ReportFile(stionDayReport)
}
