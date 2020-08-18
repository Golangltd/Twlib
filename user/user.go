package twlib_user

// 玩家结构
type UserSt struct {
	RoleUid    int64
	RoleName   string
	RoleAvatar int
	RoleLev    int
	RoleSex    int
	Coin       int64         // 金币
	Diamond    int64         // 砖石
	CardList   []*CardSt     // 角色拥有的卡牌
	LatestArea string        // 上一次的最新登录的区   区的url：ip+port
	ItemList   []*ItemSt     // 角色身上的道具，包括装备
	ChannelId  int           // 渠道Id
	ServerList []*ServerList // 整个游戏的所有区列表，从上线开始  1-30  29
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

// 道具大类型，读取对应的表
const (
	ItemMaxTypeInit = iota
	ItemCommon               // ItemCommon ==  1 常规道具
	ItemEquipment            // ItemEquipment == 2 装备
	ItemCard                 // ItemCard  == 3 卡牌
)

// 道具子类型，客户端类型
const (
	ItemMinTypeInit = iota
	ICoin                // ICoin == 1   金币
	ISilver              // ISilver == 2 银币
    ITeamExp             // ITeamExp == 3 战队经验
    ICardExp             // ICardExp == 4 卡牌经验
    IEquipmentUpExp      // IEquipmentUpExp == 5 装备强化经验
    IKindnessExp         // IKindnessExp == 6 友情点
    IArenaScore          // IArenaScore == 7 竞技场积分
    ITeamCoin            // ITeamCoin == 8 工会币
    ICardBreakCoin       // ICardBreakCoin == 9 卡牌分解货币
    IConsume             // IConsume == 10 消耗品
    IRandCardSplinter    // IRandCardSplinter == 11 随机卡牌碎片
    ICardSplinter        // ICardSplinter == 12  卡牌碎片
    IRandCard            // IRandCard == 13 随机卡牌
    ICardMagic           // ICardMagic == 14 魔法师卡牌
    IEquipmentUp         // IEquipment == 15 装备强化道具
    IEquipment           // IEquipment == 16 装备
    ITreasureBox         // ITreasureBox == 17 宝箱
    ICardDrawPaper       // ICardDrawPaper == 18 抽奖卷
    IWand                // IWand == 19 魔杖
    IArtifact            // IArtifact == 20 神器
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
	FunctionId   int     // 功能Id
	ItemData []*ItemData // 道具数据
}

// 道具数据
type ItemData struct {
	Uid      int64  // 唯一ID
	ItemId   int
	ItemNum  int64 // 道具的数量
}

//角色信息
type CardSt struct {
	CardID uint64 // 卡牌唯一ID
	Level  int    // 卡牌等级
	Skills map[int]*SkillInfo
}

//---------------------------------------------------------------------------
//用户信息
type UserInfo struct {
	ID    uint64
	Level int
	Cards map[int]*CardInfo //卡牌
}

// 角色信息
type CardInfo struct {
	CardID uint64 //卡牌唯一ID
	Level  int    //卡牌等级
	Role   string //角色ID
	Skills map[int]*SkillInfo
}

// 技能
type SkillInfo struct {
	SkillID uint64
}
