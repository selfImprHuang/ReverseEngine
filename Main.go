package main

import (
	"ReverseEngine/imp"
	"ReverseEngine/reverse"
	"ReverseEngine/support"
	"ReverseEngine/util"
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var database string
var tableName string
var filepath string
var cover bool

func main() {
	//转换函数
	reverseMain()
}

func reverseMain() {
	//报错是否直接停止程序，有的时候就是说，我可能输错了，但是我不想停止程序，那就会尴尬，这边配置一个报错不停止的标识
	panicStop := flag.Bool("panicStop", false, "错误是否进行停止")
	flag.Parse()
	conti := true
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("逆向工程生成中")
	fmt.Print("请出入你需要操作的数据库：\n")
	database, _ = reader.ReadString('\n')
	database = strings.Trim(database, "\n")
	e := reverse.NewEngine()
	e.SetReverseCheck(&support.CheckPathSupport{})

	fmt.Println(runOn(conti, reader, e, *panicStop))
}

func runOn(conti bool, reader *bufio.Reader, e *reverse.DataReverseEngine, panicStop bool) string {
	defer func() {
		if panicStop == false || conti == false {
			return
		}

		if err := recover(); err != nil {
			util.ColorPrint("程序出错，因为设置panicStop为不停止状态，所以程序继续\n请开始新一轮的反向工程参数设置\n\n\n\n........\n.........\n\n", util.FontColor.Red)
			runOn(conti, reader, e, panicStop)
		}

	}()

	for {

		if conti == false {
			return "程序运行结束"
		}
		fmt.Print("请出入你需要操作的表名,该名称不区分大小写,输出即为文件名：\n")
		tableName, _ = reader.ReadString('\n')
		tableName = strings.Trim(tableName, "\n")
		fmt.Print("请出入结构体生成位置（绝对路径）：\n")
		filepath, _ = reader.ReadString('\n')
		filepath = strings.Trim(filepath, "\n")
		fmt.Print("如果文件存在是否覆盖(Y表示覆盖)：\n")
		coverS, _ := reader.ReadString('\n')
		coverS = strings.Trim(coverS, "\n")
		if coverS == "Y" || coverS == "y" {
			cover = true
		}

		e.Engineer(database, tableName, filepath, cover)
		fmt.Print("是否停止程序,输入Y表示停止：\n")
		stop, _ := reader.ReadString('\n')
		stop = strings.Trim(stop, "\n")
		if stop == "Y" || stop == "y" {
			conti = false
		}
	}
}

func runOnce() {
	e := reverse.NewEngine()
	//e.SetContentGenerate(support.FileContentSUpport2{})
	e.SetReverseCheck(&support.CheckPathSupport{})
	e.Engineer("你的数据库", "你的表名", "F:/self_project/src/ReverseEngine/filepath", true)
	//e.Engineer("你的数据库", "testreversing", "filepath", true)
	//
	//s := support.CheckSupport{}
	//fmt.Println(&s)

	tx := reflect.ValueOf(support.CheckSupport{}).Type()
	fmt.Println(tx)
	var ok imp.ReverseCheck = (*support.CheckSupport)(nil)
	fmt.Println(ok)
	//t := reflect.ValueOf(Test{}).Type()
	//fmt.Println(t)
	//v := reflect.New(t).Elem()
	//fmt.Println(v)
}
