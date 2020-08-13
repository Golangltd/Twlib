package Chapter

const (
	ChapterInit    = iota
	ChapterLock    //  ChapterLock == 1 未解锁的章节
	ChapterUnLock  //  ChapterUnLock == 2 解锁的章节
	ChapterCurrent //  ChapterCurrent ==3 目前的章节
)

// 章节相关
type ChapterSt struct {
	ChapterId    int // 章节Id
	ChapterState int // 章节状态，是否解锁等
}
