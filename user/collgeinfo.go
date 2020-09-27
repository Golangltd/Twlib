package twlib_user

//学院信息
type CollegeInfo struct {
	CollegeID  int          //学院ID (通过学院ID就能已知等级)
	MapTalents []*MapTalent //图鉴天赋列表
}
