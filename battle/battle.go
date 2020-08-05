package twlib_battle

// 战斗类型
const (
	INIT = iota
	// pve
	RoundBattle 		// RoundBattle == 1  关卡战斗
	ClimbBattle 		// ClimbBattle == 2  爬塔战斗
	CycleBattle 		// CycleBattle == 3  周期重置战斗
	CycleRandBattle 	// CycleRandBattle  == 4 周期重置随机元素爬塔战斗
	CycleRandBattle2 	// CycleRandBattle2  == 5 重置随机元素爬塔战斗

	// pvp 
   ArenaBattle          // ArenaBattle == 6 竞技场战斗
   JacobLadderBattle    // JacobLadderBattle == 7 天梯战斗
   RaceBattle           // RaceBattle == 8 淘汰战斗
)

