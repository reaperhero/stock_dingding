package web

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/reaperhero/stock_dingding/cmd/server"
	"github.com/reaperhero/stock_dingding/model/repository"
	"github.com/reaperhero/stock_dingding/utils"
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

func GetRankBySubordinate(context echo.Context) error {
	subordinate := context.QueryParam("subordinate")
	list, err := repository.Repository.GetRankSubordinate(subordinate)
	if err != nil {
		log.Error(err)
	}

	server.CustomSortStock(list, server.SortWithSubordinateIncrease)
	var result []mStcok
	for _, stock := range list {
		l := new(mStcok)
		err := utils.StructAssignAonsistent(l, stock)
		if err != nil {
			return err
		}
		result = append(result, *l)
	}
	return context.JSON(200, map[string]interface{}{
		"count": len(result),
		"item":  result,
	})
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

		data := fmt.Sprintf("%s \n\n\n\n\n %s",
			server.EchoStock(stocks, server.SortWithSubordinateMarkValue),
			server.EchoStock(stocks, server.SortWithSubordinatePe),
		)

		err = ioutil.WriteFile("docs/subordinate/"+t+".txt", []byte(data), 0644)
		if err != nil {
			return err
		}
	}
	return context.JSON(200, "ok")
}
