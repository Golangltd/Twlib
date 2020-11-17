package modules

import . "github.com/Golangltd/Twlib/user"

/*学院表结构*/
type CollegeTable struct {
	RoleUID     int64
	CollegeID   int   //学院ID (通过学院ID就能已知等级)
	CollegeUID  int64 //学院唯一ID
	CollegeType int   //学院类型(学院一共4中类型,一般值为1,2,3,4)
}

func NewCollegeTable(roleUID int64, info *CollegeInfo) *CollegeTable {
	m := &CollegeTable{}
	m.RoleUID = roleUID
	m.CollegeID = info.CollegeID
	m.CollegeUID = info.CollegeUID
	m.CollegeType = info.CollegeType
	return m
}
