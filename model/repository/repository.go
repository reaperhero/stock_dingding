package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
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

func CreateRepository()  {
	mysqlDB, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dongfang_stock")
	if err != nil {
		log.Fatalln(err)
	}
	gormDB, _ := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dongfang_stock?charset=utf8&parseTime=True&loc=Local")
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
