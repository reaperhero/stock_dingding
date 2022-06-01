package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/reaperhero/stock_dingding/cmd/config"
	"github.com/reaperhero/stock_dingding/model"
	"log"
)

var (
	Repository *repository
)

type repository struct {
	mysqlDB *sqlx.DB
	gormDB  *gorm.DB
}

func InitRepository() {
	mysqlDB, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.IP,
		config.Config.MySQL.Port,
		config.Config.MySQL.Database,
	))
	if err != nil {
		log.Fatalln(err)
	}
	gormDB, _ := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.MySQL.User,
		config.Config.MySQL.Password,
		config.Config.MySQL.IP,
		config.Config.MySQL.Port,
		config.Config.MySQL.Database,
	))

	repo := &repository{
		mysqlDB,
		gormDB,
	}
	repo.initGormDB()
	Repository = repo
}

func (r *repository) initGormDB() {
	r.gormDB.SingularTable(true)
	r.gormDB.AutoMigrate(&model.StockPriceRanking{})
}
