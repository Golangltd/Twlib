package twlib_user

//学院信息
type CollegeInfo struct {
	CollegeID   int          //学院ID (通过学院ID就能已知等级)
	CollegeUID  int64        //学院唯一ID
	CollegeType int          //学院类型(学院一共4中类型,一般值为1,2,3,4)
	MapTalents  []*MapTalent //图鉴天赋列表
}

//排序图鉴天赋
func (m *CollegeInfo) SortMapTalent() {
	if len(m.MapTalents) <= 1 { //图鉴个数小于等于1直接返回
		return
	}
	for i := 0; i < len(m.MapTalents)-1; i++ {
		for j := i + 1; j > 0; j++ {
			if m.MapTalents[j].TalentLV < m.MapTalents[j-1].TalentLV ||
				(m.MapTalents[j].TalentLV == m.MapTalents[j-1].TalentLV && m.MapTalents[j].Quality > m.MapTalents[j-1].Quality) ||
				(m.MapTalents[j].TalentLV == m.MapTalents[j-1].TalentLV && m.MapTalents[j].Quality == m.MapTalents[j-1].Quality && m.MapTalents[j].MapID < m.MapTalents[j-1].MapID) { //天赋等级小的排在前面
				m.MapTalents[j], m.MapTalents[j-1] = m.MapTalents[j-1], m.MapTalents[j]
			} else {
				break
			}
		}
	}
}
