package web

import (
	"github.com/labstack/echo"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"github.com/reaperhero/stock_dingding/model/repository"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func health(context echo.Context) error {
	return context.JSON(200, "ok")
}

func GetStockBySubordinate(context echo.Context) error {
	list, err := repository.Repository.GetAllSubordinate()
	if err != nil {
		log.Error(err)
	}
	return context.JSON(200, list)
}

func SyncSubordinateToFile(context echo.Context) error {
	ts, err := repository.Repository.GetAllSubordinate()
	if err != nil {
		log.Errorf("[repository.Repository.GetAllSubordinate] %v", err)
		return err
	}
	for _, t := range ts {
		stocks, err := repository.Repository.GetAllStockBySubordinate(t)
		if err != nil {
			log.Errorf("[repository.Repository.GetAllSubordinate] %v", err)
			return err
		}

		data := server.EchoStock(stocks, server.SortWithSubordinateSixDaysChange)
		err = ioutil.WriteFile("docs/"+t+".md", []byte(data), 0644)
		if err != nil {
			return err
		}
	}
	return context.JSON(200, "ok")
}
