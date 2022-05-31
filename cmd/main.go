package main

import (
	_ "github.com/reaperhero/stock_dingding/cmd/init"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"github.com/reaperhero/stock_dingding/model/repository"
)

func main() {
	repository.CreateRepository()
	server.ImportExcelToDB()
}
