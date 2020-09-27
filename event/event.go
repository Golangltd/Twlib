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
	ConditionTypeRechargeSum      = 10001 // 累计充值金额
	ConditionTypeRechargeSpecify  = 10002 // 指定充值ID
	ConditionTypeLevelUser        = 10003 // 玩家等级
	ConditionTypeLevelVip         = 10004 // VIP等级
	ConditionTypeLoginContinueDay = 10005 // 连续登录天数
	ConditionTypeLoginSumDay      = 10006 // 累计登录天数
	ConditionTypeRechargeSumDay   = 10007 // 累计充值天数
	ConditionTypeConsumeSum       = 10008 // 累计消费钻石数量
)

// 初始玩家活动时候需要的玩家数据
type UserEventInitData struct {
	RoleLev int // 玩家等级
	ViPLev  int // 玩家vip等级
}

// 玩家的活动数据结构体
type UserEventSt struct {
	EventId     uint64         // 活动ID
	ConfigId    int            // 活动配置档ID
	LoopNum     int            // 循环类活动批次号，非循环类活动默认0
	CompleteNum int            // 完成次数
	BeginTime   int64          // 开启时间(开启类型-相对创角时间类型有用)
	EndTime     int64          // 结束时间(开启类型-相对创角时间类型有用)
	CloseTime   int64          // 关闭时间(开启类型-相对创角时间类型有用)
	BGetReward  bool           // 是否领取了奖励
	Conditions  []*ConditionSt // 活动条件
}

// 玩家的活动条件数据结构体
type ConditionSt struct {
	Type       int         // 条件类型
	TypeParm   int         // 条件参数
	Target     int         // 条件达成值
	TargetParm interface{} // 达成值辅助参数(如连续登录天数条件，该值存储的是玩家上次登录的时间戳)
}
