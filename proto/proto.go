package twlib_proto

// 主协议 == 规则
const (
	INIT             = iota //  INIT == 0
	GameDataProto           //  GameDataProto == 1      游戏代理的主协议       proxy server 协议
	GameDataDBProto         //  GameDataDBProto == 2    游戏的DB的主协议       db server 协议
	GameNetProto            //  GameNetProto == 3       游戏的NET主协议        心跳等
	GErrorProto             //  GErrorProto == 4        游戏的错误处理         统一错误处理
	GGateWayProto           //  GGateWayProto == 5      网关协议              暂时不用
	GGameHallProto          //  GGameHallProto == 6     游戏主场景协议         game server 协议
	GGameLoginProto         //  GGameLoginProto == 7    登录服务器协议         登录服务器主协议
	GGameGMProto            //  GGameGMProto == 8       游戏GM管理系统
	GGamePayProto           //  GGamePayProto == 9      游戏支付系统
	GGameBattleProto        //  GGameBattleProto  == 10 游戏战斗系统
	GGameConfigProto        //  GGameConfigProto  == 11 获取游戏配置  其他服务器从game游戏主逻辑服务器获取数据
	GGameCenterProto        //  GGameCenterProto  == 12 中心服务器系统
)

type GameProto64 struct {
	Protocol  int
	Protocol2 int
	ProtoVer  string // 协议版本
}

// 错误处理
type GameError struct {
	Protocol  int
	Protocol2 int
	ErrorCode int
}
