package twlib_const

// 几率
const RATE = 10000

// 离散表
const (
	LiSanTableInit          = iota
	GJ_PaoTuGuaiWu          //  GJ_PaoTuGuaiWu == 1 跑图怪物数量
	GJ_PaoTuShuaXin         //  GJ_PaoTuShuaXin == 2 跑图怪物刷新时间
	GJ_BaoXiangShouYi       //  GJ_BaoXiangShouYi == 3 挂机收益宝箱上限
	GJ_BaoXiangZhuang       //  GJ_BaoXiangShouYi == 4 挂机宝箱状态对应时间
	KSGJ_DanCiShouYi        //  KSGJ_DanCiShouYi == 5  快速挂机收益时长
	GJ_ShouYiDuiYing        //  GJ_ShouYiDuiYing == 6  挂机次数对应消耗(货币类型，数量)，次数按照配置顺序
	DJ_BeiBao               //  DJ_BeiBao == 7  道具背包上限
	ZB_BaiBao               //  ZB_BaiBao == 8 装备背包上限
	DJ_BeiBaoAllNum         //  DJ_BeiBaoAllNum == 9 道具背包总容量上限
	YJ_MainListAllNumLimit  //  YJ_MainListAllNumLimit == 10  邮件列表总容量上限
	YJ_DefaultTimeOut       //  YJ_DefaultTimeOut == 11 邮件默认有效期
	GJ_EveryDayZBNumLimit   //  GJ_EveryDayZBNumLimit == 12  每日获取装备上限(每个人)；每日的重置时间：
	BB_DefaultBagNum        //  BB_DefaultBagNum == 13 默认背包格子数量
	BB_BuyTableCost         //  BB_BuyTableCost == 14 付费扩充数量
	ZJ_WaBaoDuTiaoTime      //  ZJ_WaBaoDuTiaoTime == 15 章节挖宝的进度条的时间
	ZS_NewPlayerGetCard     //  ZS_NewPlayerGetCard == 16 新玩家赠送的卡牌
	GJ_EveryDayReJiSuanTime //  GJ_EveryDayReJiSuanTime == 17 单次挂机的重置时间
)

// 信息提示类型
const (
	NotifyTypeTips  = 1 // 飘字
	NotifyTypePanel = 2 // 面板
)
