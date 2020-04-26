package imp

import "database/sql"

type DbConnection interface {
	/*
		数据库连接接口
	*/
	CreateDbConnection(dbPath string, maxIdleConns int, connMaxLifeTime int) *sql.DB
}
