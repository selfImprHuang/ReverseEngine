
## 查询命令记录
```
查询数据库表的信息sql,这边主要是用来查询表的信息：
    select * from information_schema.TABLES where TABLE_SCHEMA='[数据库名]'

查询数据库表的字段元数据信息sql:
    show full columns from [数据库名].[表名]
查询结果：
        Field(字段名) Type(字段类型) Collation NULL(能否为空) Key(是否主键) Default(默认值) Extra Privilege Comment(注释)
```
## 使用ini配置文件读取工具
```
读取ini配置文件的工具：https://ini.unknwon.io/docs/intro/getting_started

配置信息保存路径：reverse/config.ini
支持三类配置：
    [database] //数据库配置
    [common] //通用配置
    [type_mapping] //类型映射
按照对应格式进行配置,如果没有找到对应配置则使用数据库通用配置
```

## 可拓展支持
```
可拓展支持三个接口
    ContentGenerate：文件内容拼接
    DbConnection：   数据库支持
    ReverseCheck：   校验支持

在Engine中通过以下三个方法实现配置
    func SetContentGenerate
    func SetDbConnection
    func SetReverseCheck
```

## 方法调用说明
```
    /**
     *  传入表名和生成路径,数据库名和是否覆盖按照配置文件获取
     */
    SimpleEngineer(tableName string, path string)

    /**
     *  传入表名、生成路径和是否覆盖,数据库名读取配置
     */
    func (r *DataReverseEngine) SimEngineer(tableName string, path string, cover bool)
    
     /**
     *  传入表名、生成路径和是否覆盖、数据库名
     */
    func (*DataReverseEngine) Engineer(dbName string, tableName string, path string, cover bool)
```

## 配置文件例子
```
    [database] //数据库配置
        username        = 我的数据库账号
        password        = 我的数据库密码
        tcp             = @tcp
        ip              = 我的数据库机器Ip
        port            = 3306
        database        = 我的数据库名
        charset          = utf8mb4&parseTime=True&loc=Local
        connMaxLifeTime = 100
        maxIdleConns    = 10
    
    [common] //通用配置
        def    = string //默认的映射类型
        tagKey = xorm   //tag的key
        suffix = .go    //文件后缀
        cover  = true   //存在文件是否进行覆盖，只有配置为false的时候不会覆盖，其他默认为true
    
    
    [type_mapping] //类型映射
        double = float64
        float = float64
        int = int
        tinyint = int
        bigint = int
        time = time.Time
        timestamp = time.Time
        date = time.Time
        dateTime = time.Time
```

## 方法调用例子
```
    //使用相对路径
    e := reverse.NewEngine()
    e.Engineer("数据库名字", "表名", "filepath", true)

    //使用绝对路径
    e := reverse.NewEngine()
	e.SetReverseCheck(&support.CheckPathSupport{})
	e.Engineer("数据库名字", "表名", "F:/self_project/src/ReverseEngine/filepath", true)
```

## 问题描述
```
   1.数据库连接是通过最原始的方式进行，这边后期可以修改成xorm或者其他的orm框架方式
   2.读写文件的方法可以做相应的修改，这边也是采用最原生的方式 
```

## 重复使用
直接运行Main.go方法，在控制台根据提示输入对应的数据库名、表名、生成位置、文件存在是否覆盖等参数来达到直接调用函数使用生成的目的
一次生成结束提示输入Y即可结束程序

启动时通过`panicStop`参数输入，控制程序报错不停止，除非在提示结束的时候输入Y.


## 使用
在config.ini中配置你的数据库信息，最好修改Engine中默认的数据库信息，因为这边的数据库信息都是我乱写的