package filepath

import (
	"time"
)

/*
 * @param
 * @return
 * @description 这个是生成的文件，没有格式化
 */
type Testreversing struct {
	t1  int       `xorm:"pk not null bigint(20)"` //你好
	t2  string    `xorm:"not null binary(255)"`   //你不好
	t3  string    `xorm:"bit(64)"`                //你是不是啥
	t4  string    `xorm:"blob"`                   //是的
	t5  string    `xorm:"char(255)"`              //
	t6  time.Time `xorm:"date"`                   //
	t7  string    `xorm:"datetime"`               //hello
	t8  float64   `xorm:"double(255,0)"`          //
	t9  float64   `xorm:"float(255,0)"`           //
	t10 int       `xorm:"int(255)"`               //
	t11 time.Time `xorm:"pk not null time"`       //
	t12 time.Time `xorm:"not null timestamp"`     //
	t13 string    `xorm:"varchar(255)"`           //
}
