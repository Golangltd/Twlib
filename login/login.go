package login

// 游戏客户端 登录返回的数据
type Data struct {
	Token string
	Url   string  // 反向代理的 连接地址
}
