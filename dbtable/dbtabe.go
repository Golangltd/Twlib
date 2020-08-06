package dbtable

// 数据库对应的表名
const (
	RoleTable = "gl_js_jsb"   // 角色表
	)

// 获取对应宏对应的表名
func GetDbTableEnum(table string)  string{
	switch table {
	case RoleTable :
      return RoleTable
	default:
		return ""
	}
	return ""
}