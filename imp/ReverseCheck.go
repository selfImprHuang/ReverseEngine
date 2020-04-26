package imp

import "database/sql"

type ReverseCheck interface {
	/*
		关于文件和文件夹的一致校验接口
		dirPath   文件夹路径
		tableName 数据库表名
		cover 是否覆盖已有文件
	*/
	CheckFileDir(dirPath string, tableName string, cover bool) bool

	/*
		校验数据库表是否存在
	*/
	CheckTableExist(dbName string, tableName string, db *sql.DB) bool
}
