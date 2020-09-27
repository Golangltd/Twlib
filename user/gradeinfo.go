package twlib_user

//评级信息
type GradeInfo struct {
	SRLevel  int //学籍等级
	SRLvText int //学籍等级关联的显示文本
}

func (m *GradeInfo) UpdateGradeInfo(updateLV int, nameText int) {
	m.SRLevel = updateLV
	m.SRLvText = nameText
}
