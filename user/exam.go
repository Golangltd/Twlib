package twlib_user

/*
定义数据库中exam(考试)表的数据结构
*/

//考试奖励表,记录用户当前领取了的考试奖励和当前的实践等级
type ExamAward struct {
	UserID          int64        //用户唯一ID
	PracticeLV      int          //当前实践等级
	CommonExamAward map[int]bool //普通考试奖励
	HighExamAward   map[int]bool //高级考试奖励
}
