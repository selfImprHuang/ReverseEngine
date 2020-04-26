package entity

type FieldMessage struct {
	FiledName  string //类型名称
	FieldType  string //类型
	OriginType string //数据库原始类型
	/*
		正常情况下tag的组合是 `TagKey:"pk/ not null/  FieldType"`
	*/
	TagKey  string //tag的key
	IsKey   bool   //是否主键
	Comment string //注释内容
	Default string //默认值信息
	CanNull bool   //是否可以为空
}
