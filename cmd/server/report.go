package server

import (
	"fmt"
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

	list := stock.GetHardenStockWithDays(7, 2)
	content := EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内2次涨停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetHardenStockWithDays(7, 3)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内3次涨停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetHardenStockWithDays(7, 4)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内4次涨停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetLastRoseStock(3,6)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"今日涨幅在3-6.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetLastHardenStock()
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"今日涨停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}


	list = stock.GetPlummetStockWithDays(7, 2)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内2次跌停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetPlummetStockWithDays(7, 3)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内3次跌停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

	list = stock.GetPlummetStockWithDays(7, 4)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"7日内4次跌停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}


	list = stock.GetLastFalltStock(-3,-6)
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"今日跌幅在3-6.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}


	list = stock.GetLastPlummetStock()
	content = EchoStock(list, SortWithSubordinateThreeDaysChange)
	if err := ioutil.WriteFile(dir+"今日跌停.txt", []byte(content), 0666); err != nil {
		log.Fatalf("[reportCareAboutStockTofile] ioutil.WriteFile %v", err)
	}

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
