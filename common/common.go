package twlib_common

//  GErrorProto == 4        游戏的错误处理         统一错误处理
const (
	ErrorInit = iota
	ErrorMessageProto2   // ErrorMessageProto2 == 1 服务器发送给游戏客户端
)

type ErrorMessage struct {
	Protocol  int
	Protocol2 int
	Code       string    // 错误码
	Desc       string    // 错误提示
}



