package twlib_event

// 活动结构体
type Event struct {
	EventId     uint64       // 活动ID
	ConfigId    int          // 活动配置档ID
	CompleteNum int          // 完成次数
	BGetReward  bool         // 是否领取了奖励
	Condition   []*Condition // 活动条件
}

// 活动条件
type Condition struct {
	Type     int // 条件类型
	TypeParm int // 条件参数
	Target   int // 达成值
}
