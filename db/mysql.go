package db

import (
	"fmt"

	"github.com/cenwj/echo-docs/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	confStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config().Database.DbUser,
		conf.Config().Database.DbPass,
		conf.Config().Database.DbHost,
		conf.Config().Database.DbPort,
		conf.Config().Database.DbName)

	db, err = gorm.Open(conf.Config().Database.DbType, confStr)

	//defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
}

func MysqlConn() *gorm.DB {
	return db
}
