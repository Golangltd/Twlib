package twlib_event

// 活动条件类型
const (
	ConditionTypeRechargeSingle   = 10000 // 单笔充值金额
	ConditionTypeRechargeSum      = 10002 // 累计充值金额
	ConditionTypeRechargeSpecify  = 10003 // 指定充值ID
	ConditionTypeLevelUser        = 10004 // 玩家等级
	ConditionTypeLevelVip         = 10005 // VIP等级
	ConditionTypeLoginContinueDay = 10006 // 连续登录天数
	ConditionTypeLoginSumDay      = 10007 // 累计登录天数
	ConditionTypeRechargeSumDay   = 10008 // 累计充值天数
	ConditionTypeConsumeSum       = 10009 // 累计消费钻石数量
)

// 玩家活动数据结构体
type UserEventSt struct {
	EventId     uint64         // 活动ID
	ConfigId    int            // 活动配置档ID
	CompleteNum int            // 完成次数
	BGetReward  bool           // 是否领取了奖励
	Conditions   []*ConditionSt // 活动条件
}

// 活动条件数据结构体
type ConditionSt struct {
	Type     int // 条件类型
	TypeParm int // 条件参数
	Target   int // 达成值
}

// 活动奖励
type RewardSt struct {
	ItemId   int   // 道具ID
	ItemType int   // 道具类型
	ItemNum  int64 // 道具数量
}
