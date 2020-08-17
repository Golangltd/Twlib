package twlib_server

// 代理服务器地址
const ProxyIp string = "127.0.0.1:8888"

// 所有的内部服务器的ServerId
const (
	SERVER = iota
	ProxyServerId   // ProxyServerId == 1
	GameServerId    // GameServerId   == 2
	BattleServerId  // BattleServerId == 3
	GMServerId      // GMServerId == 4
	DBServerId      // DBServerId == 5
	CenterServerId  // CenterServerId == 6
)