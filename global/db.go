package global

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lontten/lcore/lcutils"
	"github.com/lontten/ldb"
	"log"
	"os"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var DB ldb.Engine

func init() {
	conf := ldb.MysqlConf{
		Host:     os.Getenv("LDB_MYSQL_HOST"),
		Port:     os.Getenv("LDB_MYSQL_PORT"),
		DbName:   "ltask",
		User:     os.Getenv("LDB_MYSQL_USER"),
		Password: os.Getenv("LDB_MYSQL_PWD"),
		Version:  ldb.MysqlVersion5,
	}
	lcutils.LogJson(conf)

	path := "./log/ltask.log"
	writer, _ := rotatelogs.New(
		path+".%Y-%m-%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(365*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	newLogger := log.New(writer, "\r\n", log.LstdFlags)

	poolConf := ldb.PoolConf{
		MaxIdleCount: 10,
		MaxOpen:      200,
		MaxLifetime:  time.Hour,
		Logger:       newLogger,
	}
	db, err := ldb.Connect(conf, &poolConf)
	if err != nil {
		panic(err)
	}
	DB = db
}
