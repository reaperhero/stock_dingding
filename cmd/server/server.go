package server

import "github.com/reaperhero/stock_dingding/model/spider"

func SpilderRun()  {
	s := spider.NewSpider()
	s.BusinessTransaction()
}
