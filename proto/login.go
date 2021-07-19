package twlib_proto

type C2SUserLogin struct {
	AccountName string
	AccountPw   string
	DeviceId    string
	APP         string // 游戏名字
	Flag        string // 渠道名字
}
