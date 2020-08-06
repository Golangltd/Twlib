package twlib_user

// 玩家结构
type UserSt struct {
	RoleUid    int64
	RoleName   string
	RoleLev    int
	RoleSex    int
	Coin       int64         // 金币
	Diamond    int64         // 砖石
	CardList   []*CardSt     // 角色拥有的卡牌
	LatestArea int           // 上一次的最新登录的区
	ItemList   []*ItemSt     // 角色身上的道具，包括装备
	ChannelId  int           // 渠道Id
	ServerList []*ServerList // 整个游戏的所有区列表，从上线开始
}

// 游戏区的列表的状态
const (
	SDefault = iota
	SIdle    = 1 //空闲
	SBusy    = 2 //繁忙
	SNew     = 3 //新服
	SFull    = 4 //爆满
	SLast    = 5 //最近登录的
)

// 游戏服务器列表
type ServerList struct {
	ServerId   int
	ServerName string
	Url        string
	State      int
}

// 道具
type ItemSt struct {
	ItemId  int
	ItemNum int64 // 道具的数量
}

//角色信息
type CardSt struct {
	CardID uint64 //卡牌唯一ID
	Level  int    //卡牌等级
	Skills map[int]*SkillInfo
}

//---------------------------------------------------------------------------
//用户信息
type UserInfo struct {
	ID    uint64
	Level int
	Cards map[int]*CardInfo //卡牌
}

//角色信息
type CardInfo struct {
	CardID uint64 //卡牌唯一ID
	Level  int    //卡牌等级
	Role   string //角色ID
	Skills map[int]*SkillInfo
}

//技能
type SkillInfo struct {
	SkillID uint64
}
