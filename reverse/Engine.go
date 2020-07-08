package reverse

import (
	"ReverseEngine/entity"
	"ReverseEngine/imp"
	"ReverseEngine/static"
	"ReverseEngine/support"
	"ReverseEngine/util"
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

var (
	//数据库到实体映射 可在config.ini中配置
	tm = make(map[string]string)

	//通用配置 可在config.ini中配置
	def    = "string" //默认的映射类型
	tagKey = "xorm"   //tag的key
	suffix = ".go"    //文件后缀
	cover  = true     //存在文件是否进行覆盖，只有配置为false的时候不会覆盖，其他默认为true

	//数据库配置 可在config.ini中配置 格式是：”用户名:密码@tcp(IP:端口)/数据库?charset=utf8”
	username        = "你的mysql连接账号"
	password        = "你的连接密码"
	tcp             = "@tcp"
	ip              = "你的地址"
	port            = "3306"
	database        = "你的数据库" //这个是我的数据库
	charset         = "utf8mb4&parseTime=True&loc=Local"
	connMaxLifeTime = 100
	maxIdleConns    = 10

	//校验、数据库连接实现类
	cg  imp.ContentGenerate
	dbc imp.DbConnection
	rc  imp.ReverseCheck
)

//---------------------------------------初始化,策略配置,规则校验-----------------------------------------

type DataReverseEngine struct {
}

func NewEngine() *DataReverseEngine {
	return &DataReverseEngine{}
}

//创建初始化 --加载配置信息等
func init() {
	cfg := loadIniCfg() //获取ini配置文件的对象
	if nil == cfg {
		panic("配置文件 config.ini配置错误，请检查")
		return
	}
	loadCommonConfig(cfg)              //加载通用配置
	loadDataBaseConfig(cfg)            //加载数据库配置
	loadTypeMappingConfig(cfg)         //加载类型映射配置
	rc = &support.CheckSupport{}       //校验支持
	dbc = &support.DbSupport{}         //数据库支持
	cg = &support.FileContentSupport{} //文本拼接支持

}

/*
	重写自己的文件类型定义实现
*/
func (*DataReverseEngine) SetContentGenerate(face imp.ContentGenerate) {
	cg = face
}

/*
	重写自己的数据库连接工具
*/
func (*DataReverseEngine) SetDbConnection(face imp.DbConnection) {
	dbc = face
}

/*
	重写自己的校验规则
*/
func (*DataReverseEngine) SetReverseCheck(face imp.ReverseCheck) {
	rc = face
}

func checkError(err error, s string) {
	if err != nil {
		panic(s + err.Error())
	}
}

//----------------------------------逆向工程---------------------------------

/*
	数据库逆向工程
	tableName:表名
	path:创建目录的相对路径(右键选中文件夹，点击relative path得到的路径)
*/
func (r *DataReverseEngine) SimpleEngineer(tableName string, path string) {
	r.Engineer(database, tableName, path, cover)
}

/*
	数据库逆向工程
	tableName:表名
	path:创建目录的相对路径(右键选中文件夹，点击relative path得到的路径)
	cover:如果已经存在改文件是否进行覆盖
*/
func (r *DataReverseEngine) SimEngineer(tableName string, path string, cover bool) {
	r.Engineer(database, tableName, path, cover)
}

/*
	数据库逆向工程
	dbName：数据库名称
	tableName:表名
	path:创建目录的相对路径(右键选中文件夹，点击relative path得到的路径)
	cover:如果已经存在改文件是否进行覆盖
*/
func (*DataReverseEngine) Engineer(dbName string, tableName string, path string, cover bool) {
	//tableName = strings.Title(tableName) //tableName进行首字母大写
	if !rc.CheckFileDir(path, tableName, cover) { //进行通用校验
		return
	}
	dbPath := strings.Join([]string{username, ":", password, tcp, "(", ip, ":", port, ")/", database, "?", "charset=", charset}, "")
	db := dbc.CreateDbConnection(dbPath, maxIdleConns, connMaxLifeTime) //连接数据库--可配置连接不同的数据库
	if db == nil {
		log.Println("数据库连接失败，程序终止")
		return
	}
	//校验对应的数据库表是否存在
	if rc.CheckTableExist(dbName, tableName, db) {
		dbReverse(path, tableName, dbName, db)
		return
	}
	log.Println("数据库表不存在:", tableName) //不存在的话打印日志返回
}

func dbReverse(path string, tableName string, dbName string, db *sql.DB) {
	file, result := util.CreateNeedOpenFile(util.GenerateFilePath(path, tableName, static.Splice, suffix))
	if !result {
		return
	}
	cms := util.FindColumnMessage(dbName, tableName, db)         //查询数据库表字段信息
	tableComment := util.FindTableComment(dbName, tableName, db) //查询数据库表字段信息
	dErr := db.Close()                                           //关闭数据库连接
	checkError(dErr, "数据库关闭失败：")
	if nil == cms {
		log.Println("没有查到数据库字段信息略过")
	}
	fms, hasTime := buildFieldMessage(cms)                                                 //创建字段名和类型的映射
	w := bufio.NewWriter(file)                                                             //进行文件的操作
	content := cg.GenerateFileContent(path, tableName, fms, hasTime, tagKey, tableComment) //拼go文件
	_, err := w.WriteString(content)
	checkError(err, "写入出错了：")
	fErr, cErr := w.Flush(), file.Close()

	checkError(fErr, "文件写入出错了,flush错误:")
	checkError(cErr, "文件写入出错了,close错误:")
}

func buildFieldMessage(cms []entity.ColumnMessage) ([]entity.FieldMessage, bool) {
	var fms []entity.FieldMessage
	hasTime := false

	for _, cm := range cms {
		fm := &entity.FieldMessage{
			FiledName:  cm.Field,                       //字段名称
			FieldType:  typeMapping(cm.Type, &hasTime), //字段类型,进行映射
			OriginType: cm.Type,                        //数据库原始类型
			IsKey:      cm.IsKey(),                     //是否主键
			TagKey:     tagKey,                         //tag的key值，这边先写死，正常应该是可配
			Comment:    cm.GetComment(),                //注释信息
			Default:    cm.GetDefault(),                //默认值信息
			CanNull:    cm.CanNull(),                   //是否可以为空
		}

		fms = append(fms, *fm)
	}

	return fms, hasTime
}

func typeMapping(t string, has *bool) string {
	rt := strings.Split(t, "(")[0] //需要把类型后面的([长度])去掉
	value, ok := tm[rt]
	if ok {
		if value == static.TimeT {
			*has = true
		}
		return value
	}

	return def
}

//-----------------------------------读取和属性设置-----------------------------
func loadIniCfg() *ini.File {
	cfg, err := ini.Load(static.IniAddress)
	if err != nil {
		log.Println("读取失败使用原始配置", err)
		os.Exit(1)
		return nil
	}
	return cfg
}

func loadDataBaseConfig(cfg *ini.File) {
	if len(cfg.Section(static.DataBase).Keys()) == 0 {
		log.Println("配置区间为空，采用数据库配置通用配置")
		return
	}
	util.GenerateConfig("username", username, cfg.Section(static.DataBase))
	util.GenerateConfig("password", password, cfg.Section(static.DataBase))
	util.GenerateConfig("tcp", tcp, cfg.Section(static.DataBase))
	util.GenerateConfig("ip", ip, cfg.Section(static.DataBase))
	util.GenerateConfig("port", port, cfg.Section(static.DataBase))
	util.GenerateConfig("database", database, cfg.Section(static.DataBase))
	util.GenerateConfig("charset", charset, cfg.Section(static.DataBase))
	util.GenerateConfigInt("connMaxLifeTime", connMaxLifeTime, cfg.Section(static.DataBase))
	util.GenerateConfigInt("maxIdleConns", maxIdleConns, cfg.Section(static.DataBase))

}

func loadCommonConfig(cfg *ini.File) {
	if len(cfg.Section(static.Common).Keys()) == 0 {
		log.Println("配置区间为空，采用通用配置")
		return
	}
	util.GenerateConfig("def", def, cfg.Section(static.Common))          //默认的映射类型
	util.GenerateConfig("tagKey", tagKey, cfg.Section(static.Common))    //tag的key
	util.GenerateConfig("suffix", suffix, cfg.Section(static.Common))    //文件后缀
	util.GenerateConfigBool("cover", &cover, cfg.Section(static.Common)) //存在文件是否进行覆盖，只有配置为false的时候不会覆盖，其他默认为true
}

func defaultTypeMappingConfig() {
	tm["double"] = "float64"
	tm["float"] = "float64"
	tm["int"] = "int"
	tm["tinyint"] = "int"
	tm["bigint"] = "int"
	tm["time"] = "time.Time"
	tm["timestamp"] = "time.Time"
	tm["date"] = "time.Time"
	tm["dateTime"] = "time.Time"
}

func loadTypeMappingConfig(cfg *ini.File) {
	if len(cfg.Section(static.MapType).Keys()) == 0 {
		defaultTypeMappingConfig()
		log.Println("配置区间为空，采用类型映射通用配置")
		return
	}

	for _, k := range cfg.Section(static.MapType).Keys() {
		if k.Name() != "" {
			tm[k.Name()] = k.Value()
		}
	}
}
