package util

import (
	"ReverseEngine/entity"
	"database/sql"
	"fmt"
	"strings"
)

type DbUtil struct {
}

func FindColumnMessage(dbName string, tableName string, db *sql.DB) []entity.ColumnMessage {
	var cms []entity.ColumnMessage

	row, err := db.Query(strings.Join([]string{"show full columns from ", dbName, ".", tableName}, ""))
	if err != nil {
		panic("数据库查询出错：" + err.Error())
	}

	for row.Next() {
		var cm entity.ColumnMessage
		err := row.Scan(&cm.Field, &cm.Type, &cm.Collation, &cm.Null, &cm.Key, &cm.Default, &cm.Extra, &cm.Privileges, &cm.Comment)
		if err != nil {
			panic("字段赋值错误了" + err.Error())
			return nil
		}

		cms = append(cms, cm)
	}

	return cms
}

func FindTableComment(dbName string, tableName string, db *sql.DB) string {

	sql := fmt.Sprintf("select table_comment from information_schema.tables where table_schema = '%s' and table_name ='%s'", dbName, tableName)
	row, err := db.Query(sql)
	if err != nil {
		panic("数据库查询出错：" + err.Error())
	}
	fmt.Println(row)
	var s string
	for row.Next() {
		err := row.Scan(&s)
		if err != nil {
			panic("字段赋值错误了" + err.Error())
			return ""
		}
		break
	}

	return s
}
