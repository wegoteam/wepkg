package mysql

import (
	"fmt"
	"github.com/wegoteam/wepkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

const (
	MySQL = "mysql"
)

var (
	MysqlDB *gorm.DB
	once    sync.Once
)

// init
// @Description: 初始化配置
func init() {
	once.Do(func() {
		initMysqlConfig()
	})
}

// initMysqlConfig
// @Description: 初始化MySQL配置
func initMysqlConfig() {
	var mysqlConfig = &config.MySQL{}
	c := config.GetConfig()
	c.Load(MySQL, mysqlConfig)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db, mysqlConfig.Charset)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		fmt.Errorf("mysql connect failed, err: %v", err)
	}
}
