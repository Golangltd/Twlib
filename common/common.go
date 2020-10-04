package twlib_common

//  GErrorProto == 4        游戏的错误处理         统一错误处理
const (
	ErrorInit          = iota
	ErrorMessageProto2 // ErrorMessageProto2 == 1 服务器发送给游戏客户端
	ErrorLogProto2     // ErrorLogProto2 == 2 客户端发给日志服务器，日志服务器记录
)

//--------------------------------------------------------------------------------------------
// ErrorLogProto2 == 2 客户端发给日志服务器，日志服务器记录
type ErrorLog struct {
	Protocol  int // 主协议 4
	Protocol2 int
	CodeInfo  string // 代码文件名称，行数等
}

//--------------------------------------------------------------------------------------------
// ErrorMessageProto2 == 1 服务器发送给游戏客户端
type ErrorMessage struct {
	Protocol  int
	Protocol2 int
	Code      string // 错误码
	Desc      string // 错误提示
	Type      int    // 提示类型
}
