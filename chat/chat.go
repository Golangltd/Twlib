package twlib_chat

// 聊天类型
const (
	chatInit        = iota
	WorldChat       // 世界频道  WorldChat == 1
	UnionChat       // 工会频道  UnionChat == 2
	SecretChat      // 私聊     SecretChat == 3
	OtherServerChat // 跨服     OtherServerChat == 4
)
