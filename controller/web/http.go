package web

import (
	"github.com/labstack/echo"
)

func SetHttpHander(e *echo.Echo) {
	e.Static("/", "public")
	apiV1 := e.Group("/api/v1")
	{
		apiV1.GET("/health", health)
		apiV1.GET("/get-stock-subordinate", GetStockBySubordinate)
		apiV1.GET("/sync-stock-subordinate", SyncSubordinateToFile)
	}
}
