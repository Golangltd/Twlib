package twlib_user

//图鉴天赋
type MapTalent struct {
	MapID      int   //图鉴ID
	MapUID     int64 //图鉴唯一ID
	TalentLV   int   //天赋等级
	Quality    int   //品质
	IsGraduate bool  //课程是否毕业
}
