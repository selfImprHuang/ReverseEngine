/*
 *  @Author : huangzj
 *  @Time : 2020/4/17 14:41
 *  @Description：
 */

package util

import (
	"gopkg.in/ini.v1"
	"strconv"
)

func GenerateConfigBool(key string, real *bool, section *ini.Section) {
	if section.Key(key).String() != "false" {
		*real = true
		return
	}
	*real = false
}

func GenerateConfig(key string, real string, section *ini.Section) {
	if section.Key(key).String() != "" {
		real = section.Key(key).String()
	}
}

func GenerateConfigInt(key string, real int, section *ini.Section) {
	if section.Key(key).String() != "" {
		value, err := strconv.Atoi(section.Key(key).String())
		if err == nil {
			real = value
		}
	}
}
