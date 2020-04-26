package imp

import "ReverseEngine/entity"

type ContentGenerate interface {
	/*
		内容制造接口
	*/
	GenerateFileContent(filePath string, fileName string, fms []entity.FieldMessage, hasTime bool, tagKey string) string
}
