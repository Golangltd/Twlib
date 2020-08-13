package Twlib_Chapter

const (
	RoundInit    = iota
	RoundLock    //  RoundLock == 1 未解锁的关卡
	RoundUnLock  //  RoundUnLock == 2 解锁的关卡
	RoundCurrent //  RoundCurrent ==3 目前的关卡
)

// 关卡相关
type RoundSt struct {
	RoundId    int // 关卡Id
	RoundState int // 关卡状态，是否解锁等
}

