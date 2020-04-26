package support

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DbSupport struct{}

func (dbc *DbSupport) CreateDbConnection(dbPath string, maxIdleConns int, connMaxLifeTime int) *sql.DB {
	var db *sql.DB
	db, _ = sql.Open("mysql", dbPath)

	//设置数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(connMaxLifeTime))
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(maxIdleConns)

	if err := db.Ping(); err != nil {
		panic("数据库连接失败")
		return nil
	}

	log.Println("数据库连接成功")
	return db
}
