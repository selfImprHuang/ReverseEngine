package support

import (
	"ReverseEngine/entity"
	"ReverseEngine/static"
	"strings"
)

type FileContentSupport struct{}

func (FileContentSupport) GenerateFileContent(filePath string, fileName string, fms []entity.FieldMessage, hasTime bool, tagKey string) string {
	var build strings.Builder
	pName := splitLast(filePath, static.Splice)

	writePackage(&build, pName)        //写文件的引用包
	writeImportIfTime(&build, hasTime) //如果有time.Time类型的需要写入import
	buildStructHead(&build, fileName)  //写入结构体头

	//对于没有字段的返回空定义的文件
	if len(fms) == 0 {
		build.WriteString(static.RightBrace)
		return build.String()
	}

	//写中间字段的操作
	for _, fm := range fms {
		build.WriteString("	")
		build.WriteString(fm.FiledName)
		build.WriteString(" ")
		build.WriteString(fm.FieldType)
		build.WriteString(" ")
		build.WriteString(generateTag(fm, tagKey))
		build.WriteString(" ")
		build.WriteString(static.DoubleInclinedRod)
		build.WriteString(fm.Comment)
		build.WriteString(static.LineFeed)
	}

	build.WriteString(static.RightBrace)

	return build.String()
}

func generateTag(message entity.FieldMessage, tagKey string) string {
	var build strings.Builder
	build.WriteString("`")
	build.WriteString(tagKey)
	build.WriteString(static.Colon)
	generateTagValue(message, &build)
	build.WriteString("`")
	return build.String()
}

func generateTagValue(message entity.FieldMessage, builder *strings.Builder) {
	builder.WriteString("\"")
	if message.IsKey {
		builder.WriteString("pk ")
	}
	if !message.CanNull {
		builder.WriteString("not null ")
	}

	builder.WriteString(message.OriginType)
	builder.WriteString("\"")
}

func buildStructHead(build *strings.Builder, fileName string) {
	build.WriteString("type ")
	build.WriteString(fileName)
	build.WriteString(" struct")
	build.WriteString(static.LeftBrace)
	build.WriteString(static.LineFeed)
}

func writeImportIfTime(build *strings.Builder, hasTime bool) {
	//如果有类型匹配上时间，那么需要在头部加上import
	if hasTime {
		build.WriteString("import (")
		build.WriteString("\n	")
		build.WriteString("\"time\"")
		build.WriteString("\n")
		build.WriteString(")")
		build.WriteString(static.LineFeed)
		build.WriteString(static.LineFeed)
	}
}

func writePackage(build *strings.Builder, pName string) {
	build.WriteString("package ")
	build.WriteString(pName)
	build.WriteString(static.LineFeed)
	build.WriteString(static.LineFeed)
}

/*
	返回按照规则切割的最后一个字符串
*/
func splitLast(filePath string, split string) string {
	ss := strings.Split(filePath, split)
	return ss[len(ss)-1]
}
