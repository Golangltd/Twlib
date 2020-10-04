package twlib_user

// 玩家结构
type UserSt struct {
	RoleUid             int64                //可用作学号
	RoleName            string               //角色名字
	RoleAvatar          int                  //角色头像
	RoleLev             int                  //角色等级
	RoleSex             int                  //角色性别
	RoleExp             int                  // 巫师经验
	Coin                int64                // 金币
	Diamond             int64                // 砖石
	TotalPower          int                  // 总战力
	Association         string               // 协会
	CardList            []*CardInfo          // 角色拥有的卡牌
	LatestArea          string               // 上一次的最新登录的区   区的url：ip+port
	//ItemList            []*ItemData            // 背包里的道具
	ItemList            []*ItemData            // 背包里的道具
	EquipData           *EquipData           // 背包里的的装备道具
	ChannelId           int                  // 渠道Id
	ServerList          []*ServerList        // 整个游戏的所有区列表，从上线开始  1-30  29 数据更新操作
	ChapterInfo         *ChapterInfo         // 当前章节+当前关卡
	ClearanceDuplicates map[int]bool         // 通关副本
	CollegesInfo        map[int]*CollegeInfo // 学院信息
	GradeInfo           *GradeInfo           // 评级信息
	RegisterTime        int64                // 玩家注册角色时间
}

//updateLV: 更新的等级 nameText(学籍关联的文本)
func (m *UserSt) UpdateGradeInfo(updateLV int, nameText int) {
	m.RoleExp = 0 //升级之后经验置为0
	m.GradeInfo.UpdateGradeInfo(updateLV, nameText)
}

//power:战斗力
func (m *UserSt) UpdateTotalPower(power int) {
	m.TotalPower = power
}

//学院升级处理
//pastCollegeID:过去的学院ID upgradedCollegeID:当前的学院ID
func (m *UserSt) UpgradeHandle(pastCollegeID int, upgradedCollegeID int) {
	collegeInfo := m.CollegesInfo[pastCollegeID]
	collegeInfo.CollegeID = upgradedCollegeID
	delete(m.CollegesInfo, pastCollegeID)
	m.CollegesInfo[upgradedCollegeID] = collegeInfo
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
	ItemCommon      // ItemCommon ==  1 常规道具
	ItemEquipment   // ItemEquipment == 2 装备
	ItemCard        // ItemCard  == 3 卡牌
)

// 道具子类型，客户端类型
const (
	ItemMinTypeInit   = iota
	ICoin             // ICoin == 1   金币
	ISilver           // ISilver == 2 银币
	ITeamExp          // ITeamExp == 3 战队经验
	ICardExp          // ICardExp == 4 卡牌经验
	IEquipmentUpExp   // IEquipmentUpExp == 5 装备强化经验
	IKindnessExp      // IKindnessExp == 6 友情点
	IArenaScore       // IArenaScore == 7 竞技场积分
	ITeamCoin         // ITeamCoin == 8 工会币
	ICardBreakCoin    // ICardBreakCoin == 9 卡牌分解货币
	IConsume          // IConsume == 10 消耗品
	IRandCardSplinter // IRandCardSplinter == 11 随机卡牌碎片
	ICardSplinter     // ICardSplinter == 12  卡牌碎片
	IRandCard         // IRandCard == 13 随机卡牌
	ICardMagic        // ICardMagic == 14 魔法师卡牌
	IEquipmentUp      // IEquipment == 15 装备强化道具
	IEquipment        // IEquipment == 16 装备
	ITreasureBox      // ITreasureBox == 17 宝箱
	ICardDrawPaper    // ICardDrawPaper == 18 抽奖卷
	IWand             // IWand == 19 魔杖
	IArtifact         // IArtifact == 20 神器
	IShowMapChip      // IShowMapChip == 21 图鉴碎片
	ISQChip           // ISQChip == 22 神器碎片
	IZSChip           // IZSChip == 23 杖身碎片
	IZXMaterial       // IZXMaterial == 24 杖心材料
	IMZDrawing        // IMZDrawing == 25 魔杖图纸
)

// 游戏服务器列表
type ServerList struct {
	ServerId   int
	ServerName string
	Url        string
	State      int
	UserInfo   *UserSt // 存储到数据库
}

// 道具
type ItemSt struct {
	FunctionId int         // 功能Id
	ItemData   []*ItemData // 道具数据
}

// 道具数据
type ItemData struct {
	ItemUid   int64 // 唯一ID
	ItemId    int
	ItemType  int
	ItemNum   int64 // 道具的数量
}

// 数据库 存储道具的结构
type ItemSaveSt struct {
	AccountId  uint64
	RoleId     uint64
	FunctionId int // 功能Id
	ItemId     int
	ItemNum    int64 // 道具的数量
}

//---------------------------------------------------------------------------
//用户信息
type UserInfo struct {
	ID    uint64
	Level int
	Cards map[int]*CardInfo // 卡牌
}

// 卡牌信息结构
type CardInfo struct {
	CardUid       uint64       // 数据库唯一Id
	CardID        uint64       // 卡牌ID
	Level         int          // 卡牌等级
	Skills        []*SkillInfo // 技能列表
	Equips        []*EquipSt   // 卡牌的装备
	AttributeInfo *AttributeSt // 战斗力
	Stars         int          // 星级
	IsShow        bool         // 图鉴系统，false:没有打开过图鉴系统；true:打开过图鉴系统
}

// 属性加成
type AttributeSt struct {
	BattlePower  uint64 // 战斗力
	HPPower      uint64 // 血量
	AttackPower  uint64 // 攻击力
	DefensePower uint64 // 防御力
}

// 技能
type SkillInfo struct {
	SkillId  int64
	SkillLev int // 数据操作
	Position int // 0 1 2  -1:没有使用
}

// 玩家装备数据
type EquipData struct {
	FunctionId int        // 功能ID 固定值 101
	EquipSts   []*EquipSt // 装备列表

}

// 装备结构体
type EquipSt struct {
	UID      int64 // 唯一ID
	ConfID   int   // 配置表ID
	Star     int   // 装备星级
	Camp     int   // 属性加成的阵营
	CampRate int   // 阵营加成比例百分比
	Num      int   // 数量
	Exp      int   // 当前强化经验
	Power    int   // 评分(战力)
}
