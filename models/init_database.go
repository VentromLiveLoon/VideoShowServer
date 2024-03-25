package models

import (
	"github.com/beego/beego/logs"
	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/core/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db_driver, _ := config.String("db_driver")
	db_host, _ := config.String("db_host")
	db_port, _ := config.String("db_port")
	db_user, _ := config.String("db_user")
	db_password, _ := config.String("db_password")
	db_name, _ := config.String("db_name")
	db_charset, _ := config.String("db_charset")
	url := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?" + "charset=" + db_charset
	logs.Info(url)
	orm.RegisterDataBase("default", db_driver, url)
}
