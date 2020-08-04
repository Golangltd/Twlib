package user

//用户信息
type UserInfo struct {
	ID    uint64
	Level int
	Cards map[int]*CardInfo //卡牌
}

//角色信息
type CardInfo struct {
	CardID uint64 //卡牌唯一ID
	Level  int    //卡牌等级
	Role   string   //角色ID
	Skills map[int]*SkillInfo
}

//技能
type SkillInfo struct {
	SkillID uint64
}
