package Db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"sync"
)

type Orm interface {
	NewOrm() *xorm.Engine
	RegisterModels(model interface{})
	RegisterFixModels(model interface{})
}

type Engine struct {
	conn *xorm.Engine
	Fix  string
}

func newEngine(driver string, urlConn string) Orm {
	e := new(Engine)
	e.init(driver, urlConn)
	return e
}

func (e *Engine) init(driver string, urlConn string) {
	var err error
	e.conn, err = xorm.NewEngine(driver, urlConn)
	if err != nil {
		// 记录日志
		fmt.Println("mysql connection failed error ", err)
		return
	}
	e.conn.SetMaxOpenConns(10)
	e.conn.SetMaxIdleConns(3)
}

func (e *Engine) NewOrm() *xorm.Engine {
	return e.conn
}

func (e *Engine) RegisterModels(model interface{}) {
	e.conn.SetTableMapper(core.SameMapper{})
	if err := e.conn.Sync2(model); err != nil {
		return
	}
}

func (e *Engine) RegisterFixModels(model interface{}) {
	e.conn.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, e.Fix))
	if err := e.conn.Sync2(model); err != nil {
		return
	}
}

var (
	once   = new(sync.Once)
	DbConn Orm
)

func InitMysql(driver string, urlConn string) {
	once.Do(func() {
		DbConn = newEngine(driver, urlConn)
	})
}