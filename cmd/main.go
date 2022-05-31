package main

import (
	"fmt"
	_ "github.com/reaperhero/stock_dingding/cmd/init"
	"github.com/reaperhero/stock_dingding/model/repository"
	"github.com/reaperhero/stock_dingding/service/stock_analyse"
)

func main() {
	repository.CreateRepository()
	//server.ImportExcelToDB()
	list,maxCount := stock_analyse.DailyLimitStatistics()
	for i := maxCount; i > 0; i-- {
		for hanYe, _ := range list {
			stocks := list[hanYe]
			if len(stocks) == i {
				fmt.Println(hanYe, stocks)
			}
		}
	}
}
