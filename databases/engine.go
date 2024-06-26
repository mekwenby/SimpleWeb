package databases

/*
数据库引擎
*/

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var Engine *xorm.Engine

func init() {
	// 在应用程序初始化阶段创建数据库引擎
	Engine = CreateEngine()
}
func CreateEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "image.db")
	if err != nil {
		panic(err)
	}
	return engine
}
