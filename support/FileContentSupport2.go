package support

import (
	"ReverseEngine/entity"
	"fmt"
)

type FileContentSUpport2 struct{}

func (FileContentSUpport2) GenerateFileContent(filePath string, fileName string, fms []entity.FieldMessage, hasTime bool, tagKey string, string string) string {
	fmt.Println("12321321")
	return "       "
}
