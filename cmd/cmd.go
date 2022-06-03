package main

import (
	_ "github.com/reaperhero/stock_dingding/cmd/config"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"github.com/reaperhero/stock_dingding/model/repository"
)

func init()  {
	repository.InitRepository()
}

func main() {
	server.Execute()
}
