package entity

import "database/sql"

type ColumnMessage struct {
	Field      string         //字段名称
	Type       string         //字段类型
	Collation  sql.NullString //编码信息
	Null       sql.NullString //是否允许为空
	Key        sql.NullString //是否是主键
	Default    sql.NullString //默认值
	Extra      sql.NullString //不懂
	Privileges sql.NullString //执行权限
	Comment    sql.NullString //注释信息
}

/*
	返回结果判断是不是主键
*/
func (c *ColumnMessage) IsKey() bool {
	if c.Key.Valid && c.Key.String == "PRI" {
		return true
	}
	return false
}

/*
	返回结果是否为空信息
*/
func (c *ColumnMessage) CanNull() bool {
	if c.Null.Valid && c.Null.String == "NO" {
		return false
	}
	return true
}

/*
	返回注释信息，没有注释信息返回空字符串
*/
func (c *ColumnMessage) GetComment() string {
	if c.Comment.Valid {
		return c.Comment.String
	}

	return ""
}

/*
	返回默认值信息
*/
func (c *ColumnMessage) GetDefault() string {
	if c.Default.Valid {
		return c.Default.String
	}
	return ""
}
