package Twlib_Chapter

const (
	ChapterInit    = iota
	ChapterLock    //  ChapterLock == 1 未解锁的章节
	ChapterUnLock  //  ChapterUnLock == 2 解锁的章节
	ChapterCurrent //  ChapterCurrent ==3 目前的章节
)

// 章节相关
type ChapterSt struct {
	ChapterId    int              // 章节Id
	ChapterState int              // 章节状态，是否解锁等
	PlayerData   []*ChapterUserSt // 章节的用户信息
}

// 玩家结构
type ChapterUserSt struct {
	RoleUid    int64
	RoleName   string
	RoleAvatar int
	RoleLev    int
	RoleSex    int
	RoleExp    int   // 巫师经验
	Coin       int64 // 金币
	Diamond    int64 // 砖石
	ChannelId  int   // 渠道Id
}
