package cfgame

// 获取配置数据库的协议
const (
	CFInit              = iota
	GetCFGameDataProto2 // GetCFGameDataProto2 == 1 获取配置数据库数据
)

type GetCFGameData struct {
	Protocol  int
	Protocol2 int
	TableName string // 表名
}
