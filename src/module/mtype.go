package module

type Type string

const (
	TYPE_DOWNLOADER Type = "downloader"
	TYPE_ANALYZER   Type = "analyzer"
	TYPE_PIPELINE   Type = "pipeline"
)

var legalTypeLetterMap = map[Type]string{
	TYPE_DOWNLOADER: "D",
	TYPE_ANALYZER:   "A",
	TYPE_PIPELINE:   "P",
}

var legalLetterTypeMap = map[string]Type{
	"D": TYPE_DOWNLOADER,
	"A": TYPE_ANALYZER,
	"P": TYPE_PIPELINE,
}

func CheckType(moduleType Type, module Module) bool {
	if moduleType == "" || module == nil {
		return false
	}
	switch moduleType {
	case TYPE_DOWNLOADER:
		if _, ok := module.(Downloader); ok {
			return true
		}
	case TYPE_ANALYZER:
		if _, ok := module.(Analyzer); ok {
			return true
		}
	case TYPE_PIPELINE:
		if _, ok := module.(Pipeline); ok {
			return true
		}
	}
	return false
}

// LegalType 用于判断给定的组件类型是否合法。
func LegalType(moduleType Type) bool {
	if _, ok := legalTypeLetterMap[moduleType]; ok {
		return true
	}
	return false
}

// GetType 用于获取组件的类型。
// 若给定的组件ID不合法则第一个结果值会是false。
func GetType(mid MID) (bool, Type) {
	parts, err := SplitMID(mid)
	if err != nil {
		return false, ""
	}
	mt, ok := legalLetterTypeMap[parts[0]]
	return ok, mt
}
