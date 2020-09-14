package twlib_event

// 当前运行的活动数据结构体
type EventSt struct {
	EventId   uint64 // 活动ID
	ConfigId  int    // 配置表ID
	State     int    // 活动状态
	LoopNum   int    // 循环类活动批次号，非循环类活动默认0
	BeginTime int64  // 开启时间
	EndTime   int64  // 结束时间
	CloseTime int64  // 关闭时间
}

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

// 玩家的活动数据结构体
type UserEventSt struct {
	EventId     uint64         // 活动ID
	ConfigId    int            // 活动配置档ID
	CompleteNum int            // 完成次数
	BGetReward  bool           // 是否领取了奖励
	Conditions  []*ConditionSt // 活动条件
}

// 玩家的活动条件数据结构体
type ConditionSt struct {
	Type     int // 条件类型
	TypeParm int // 条件参数
	Target   int // 达成值
}
