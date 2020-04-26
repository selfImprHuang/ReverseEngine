package support

import (
	"ReverseEngine/entity"
	"ReverseEngine/static"
	"ReverseEngine/util"
	"database/sql"
	"log"
	"reflect"
	"strings"
)

type CheckSupport struct{}

func (*CheckSupport) CheckFileDir(path string, tableName string, cover bool) bool {
	b, realPath := pathExist(path)
	//先查询一下路径是否为文件夹
	if !b {
		log.Println("文件夹路径不存在略过")
		return false
	}
	//如果是不覆盖的情况，判断是否有这个文件，如果有的话直接返回
	if !cover && util.IsFileExist(strings.Join([]string{path, static.Splice, tableName}, "")) {
		return false
	}

	//判断读写权限
	if !util.HasReadWritePermission(realPath) {
		log.Println("没有该文件夹的读写权限：", realPath)
		return false
	}

	return true
}

func (*CheckSupport) CheckTableExist(dbName string, tableName string, db *sql.DB) bool {
	var tm entity.TableMessage
	err := db.QueryRow(strings.Join([]string{"select TABLE_NAME from information_schema.TABLES where TABLE_SCHEMA=? and TABLE_NAME=?"}, ""), dbName, tableName).Scan(&tm.TableName)
	//查询失败
	if err != nil {
		panic("数据库查询失败:" + err.Error())
		return false
	}
	//查询没有数据
	if !reflect.ValueOf(tm).IsValid() || tm.TableName == "" {
		panic("没有表数据信息：" + tableName)
		return false
	}

	return true
}

func pathExist(path string) (bool, string) {
	pp := util.GetCurrentPath()
	realPath := strings.Join([]string{pp, static.Splice, path}, "")
	return util.IsDirExist(strings.Join([]string{pp, static.Splice, path}, "")), realPath
}
