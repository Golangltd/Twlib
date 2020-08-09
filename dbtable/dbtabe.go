package twlib_dbtable

// 数据库对应的表名
const (
	RoleTable   = "gl_js_jsb"     // 角色表
	ItemTable   = "gl_dj_djb"     // 道具表
	SkillTable1 = "gl_jn_jncfb"   // 技能1表
	SkillTable2 = "gl_jn_jnewmbb" // 技能2表
	SkillTable3 = "gl_jn_jntsxgb" // 技能3表
	SkillTable4 = "gl_jn_jnzb"    // 技能4表
)

// 获取对应宏对应的表名
func GetDbTableEnum(table string) string {
	switch table {
	case RoleTable:
		return RoleTable
	case ItemTable:
		return ItemTable
	case SkillTable1:
		return SkillTable1
	case SkillTable2:
		return SkillTable2
	case SkillTable3:
		return SkillTable3
	case SkillTable4:
		return SkillTable4
	default:
		return ""
	}
	return ""
}
