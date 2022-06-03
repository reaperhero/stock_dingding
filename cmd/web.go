package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/reaperhero/stock_dingding/cmd/config"
	_ "github.com/reaperhero/stock_dingding/cmd/config"
	"github.com/reaperhero/stock_dingding/controller/web"
	"github.com/reaperhero/stock_dingding/model/repository"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	repository.InitRepository()
}

func httpRun() {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if _, ok := err.(*echo.HTTPError); !ok {
			log.WithError(err).Errorln("[echo.HTTPErrorHandler]", c.Request().Method, c.Request().URL)
		}
		e.DefaultHTTPErrorHandler(err, c)
	}

	//dataPath := os.Getenv("DATA_PATH")
	//accessLogFile, err := os.OpenFile(dataPath+"access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatalln("Failed create access log file")
	//}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{},
		AllowMethods:     []string{echo.GET, echo.OPTIONS, echo.POST},
		AllowCredentials: true,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: os.Stdout}))
	e.Use(middleware.Recover())

	web.SetHttpHander(e)
	err := e.Start(fmt.Sprintf(":%d", config.Config.Web.Port))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	httpRun()
}
