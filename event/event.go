package twlib_event

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
