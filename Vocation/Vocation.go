package twlib_Vocation

type Vocation = int //职业
// 职业
const (
	Warrior Vocation = 1 //战士
	Ranger  Vocation = 2 //游侠
	Mage    Vocation = 3 //法师
	Assist  Vocation = 4 //辅助
)


type AttackType = int //攻击类型
const (
	Physical AttackType = -1 //物理
	Magic    AttackType = 1  //魔法
)

type Position = int  //位置信息
const (
	BackUp    Position = 0 //后上
	FrontUp   Position = 1 //前上
	BackMid   Position = 2 //后中
	FrontMid  Position = 3 //前中
	BackDown  Position = 4 //后下
	FrontDown Position = 5 //前下
)
