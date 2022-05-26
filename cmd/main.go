package main

import (
	_ "github.com/reaperhero/stock_dingding/cmd/init"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"time"
)

func main() {
	server.SpilderRun()
	time.Sleep(time.Second * 100)
}
