package web

import (
	"github.com/labstack/echo"
	"github.com/reaperhero/stock_dingding/model/repository"
	log "github.com/sirupsen/logrus"
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
