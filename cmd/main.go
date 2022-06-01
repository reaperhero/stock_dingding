package main

import (
	_ "github.com/reaperhero/stock_dingding/cmd/config"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"github.com/reaperhero/stock_dingding/model/repository"


)

func main() {
	repository.InitRepository()
	server.Execute()

}
