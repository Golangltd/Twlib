package modules

import . "github.com/Golangltd/Twlib/user"

/*图鉴表*/

type MapTable struct {
	CollegeUID int64 //学院唯一ID
	MapID      int   //图鉴ID
	MapUID     int64 //图鉴唯一ID
	TalentLV   int   //天赋等级
	Quality    int   //品质
	IsGraduate bool  //课程是否毕业
}

func NewMapTable(collegeUID int64, data *MapTalent) *MapTable {
	m := &MapTable{}
	m.CollegeUID = collegeUID
	m.MapID = data.MapID
	m.MapUID = data.MapUID
	m.TalentLV = data.TalentLV
	m.Quality = data.Quality
	m.IsGraduate = data.IsGraduate
	return m
}
