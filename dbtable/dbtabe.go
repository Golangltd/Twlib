package twlib_dbtable

// 数据库对应的表名

const (

	// 道具
	ItemTable = "gl_dj_djb"        // 道具表
	gl_dj_djlxb = "gl_dj_djlxb"    // 道具类型表
	gl_dj_djzlxb = "gl_dj_djzlxb"  // 道具子类型表
	gl_dl_dlpzb = "gl_dl_dlpzb"    // 掉落配置表

	// 关卡
	gl_gk_cjpzb = "gl_gk_cjpzb"    // 暂时无
	Gl_gk_gjsjb = "gl_gk_gjsjb"    // 章节数据表
	Gl_gk_gkgwb = "gl_gk_gkgwb"    // 关卡怪物表
	Gl_gk_gkpzb = "gl_gk_gkpzb"    // 关卡配置表
	Gl_gk_gwzpzb = "gl_gk_gwzpzb"  // 关卡总配置表
	Gl_gk_gwzxb = "gl_gk_gwzxb"    // 怪物属性表
	gl_gk_zjbxpzb = "gl_gk_zjbxpzb"//
	gl_gk_zjpzb = "gl_gk_zjpzb"    //

	// 技能
	gl_gn_gnkqpzb = "gl_gn_gnkqpzb"//
	gl_gn_gntjlxb = "gl_gn_gntjlxb"//
	SkillTable1 = "gl_jn_jncfb"    // 技能触发表
	SkillTable2 = "gl_jn_jnewmbb"  // 技能额外目标
	SkillTable3 = "gl_jn_jntsxgb"  // 技能特殊效果
	SkillTable4 = "gl_jn_jnzb"     // 技能总表

	// 角色
	RoleTable = "gl_js_jsb"        // 角色表
	gl_js_jssxxs = "gl_js_jssxxs"  //
	gl_ls_lssjb = "gl_ls_lssjb"
	gl_mz_mzdzpzxsb = "gl_mz_mzdzpzxsb"
	gl_mz_mztxsxc = "gl_mz_mztxsxc"
	gl_pz_pzjjpzb = "gl_pz_pzjjpzb"
	gl_rw_rwpzb = "gl_rw_rwpzb"
	gl_sd_sdlbpzb = "gl_sd_sdlbpzb"
	gl_sd_sdsppzb = "gl_sd_sdsppzb"
	gl_sz_szpzb = "gl_sz_szpzb"
	gl_tj_kptj0bb = "gl_tj_kptj0bb"
	gl_tj_sqtjb = "gl_tj_sqtjb"
	gl_tj_yxkptjb = "gl_tj_yxkptjb"
	gl_ts_tspzb = "gl_ts_tspzb"
	gl_wb_wbb = "gl_wb_wbb"
	gl_wse_wsjyb = "gl_wse_wsjyb"
	gl_xj_xjkcb = "gl_xj_xjkcb"
	gl_xj_xjpzb = "gl_xj_xjpzb"
	gl_xm_sjxmb = "gl_xm_sjxmb"
	gl_yj_yjpzb = "gl_yj_yjpzb"
	gl_zb_zbb = "gl_zb_zbb"
	gl_zw_zwpzb = "gl_zw_zwpzb"
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
	case Gl_gk_gjsjb:
		return Gl_gk_gjsjb
	case Gl_gk_gkgwb:
		return Gl_gk_gkgwb
	case Gl_gk_gkpzb:
		return Gl_gk_gkpzb
	case Gl_gk_gwzpzb:
		return Gl_gk_gwzpzb
	case Gl_gk_gwzxb:
		return Gl_gk_gwzxb
	default:
		return ""
	}
	return ""
}