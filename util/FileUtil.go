package util

import (
	"log"
	"os"
	"strings"
)

/*
	获取当前项目的路径
*/
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/*
	判断路径是否存在并且为文件夹
*/
func IsDirExist(dirPath string) bool {
	fi, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	}

	return false
}

/*
	判断路径是否存在并且为文件
*/
func IsFileExist(filePath string) bool {
	fi, err := os.Stat(filePath)
	//文件不存在或者是文件夹的直接返回false
	if err != nil || fi.IsDir() {
		return false
	}

	return true
}

/*
	判断是否有读权限
*/
func HasReadPermission(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}

	return true

}

/*
	判断是否有写权限
*/
func HasWritePermission(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return false
		}
	}

	return true
}

/*
	判断是否有读+写权限
*/
func HasReadWritePermission(filePath string) bool {
	return HasReadPermission(filePath) && HasWritePermission(filePath)
}

/*
	创建文件
*/
func CreateFile(filePath string) (*os.File, bool) {
	file, err := os.Create(filePath)
	if err != nil {
		panic("创建文件失败" + err.Error())
		return nil, false
	}

	return file, true
}

/*
	返回拼接的文件路径,这边类的首字母需要大写
	dirPath 文件夹路径
	filePath 文件路径
	splice	路径分隔符
	suffix 后缀
*/
func GenerateFilePath(dirPath string, filePath string, splice string, suffix string) string {
	return strings.Join([]string{dirPath, splice, strings.Title(filePath), suffix}, "")
}

/*
	创建（如果需要）并已读写权限打开文件
	os.O_RDONLY // 只读
    os.O_WRONLY // 只写
    os.O_RDWR // 读写
    os.O_APPEND // 往文件中添建（Append）
    os.O_CREATE // 如果文件不存在则先创建
    os.O_TRUNC // 文件打开时裁剪文件
    os.O_EXCL // 和O_CREATE一起使用，文件不能存在
    os.O_SYNC // 以同步I/O的方式打开
*/
func CreateNeedOpenFile(filePath string) (*os.File, bool) {
	//先删除再创建
	if IsFileExist(filePath) {
		err := os.Remove(filePath)
		if err != nil {
			panic("文件删除失败：" + err.Error())
			return nil, false
		}
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("打开文件失败：" + err.Error())
		return nil, false
	}
	return file, true
}
