package twlib_user

//学院信息
type CollegeInfo struct {
	CollegeID  int          //学院ID (通过学院ID就能已知等级)
	MapTalents []*MapTalent //图鉴天赋列表
}

//排序图鉴天赋
func (m *CollegeInfo) SortMapTalent() {
	if len(m.MapTalents) <= 1 { //图鉴个数小于等于1直接返回
		return
	}
	for i := 0; i < len(m.MapTalents)-1; i++ {
		for j := i + 1; j > 0; j++ {
			if m.MapTalents[j].TalentLevel < m.MapTalents[i].TalentLevel ||
				(m.MapTalents[j].TalentLevel == m.MapTalents[i].TalentLevel && m.MapTalents[j].CurrentQuality > m.MapTalents[i].CurrentQuality) ||
				(m.MapTalents[j].TalentLevel == m.MapTalents[i].TalentLevel && m.MapTalents[j].CurrentQuality == m.MapTalents[i].CurrentQuality && m.MapTalents[j].MapID < m.MapTalents[i].MapID) { //天赋等级小的排在前面
				m.MapTalents[i], m.MapTalents[j] = m.MapTalents[j], m.MapTalents[i]
			} else {
				break
			}
		}
	}
}


