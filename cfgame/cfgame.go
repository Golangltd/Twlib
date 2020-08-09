package cfgame

// 获取配置数据库的协议
const (
	CFInit              = iota
	GetCFGameDataProto2 // GetCFGameDataProto2 == 1 获取配置数据库数据
	CFRetGameDataProto2 // CFRetGameDataProto2 == 2 返回的数据
)

type GetCFGameData struct {
	Protocol  int
	Protocol2 int
	TableName string // 表名
	ServerId  string // 服务器Id
}

type CFRetGameData struct {
	Protocol  int
	Protocol2 int
	TableName string // 表名
	Data      interface{}
}
