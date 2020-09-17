package twlib_mail

// 邮件结构体(服务器内部使用)
type MailSt struct {
	MailId 		uint64			// 邮件唯一ID
	Type 		int				// 类型
	SenderId	uint64			// 发送人ID
	Title 		string 			// 标题
	Description string 			// 内容
	SendTime 	int64 			// 发送时间戳
	ExpireTime	int64			// 过期时间戳
	BTemplate 	bool 			// 是否模板邮件
	BRead 		bool 			// 是否阅读
	BGetAtt 	bool 			// 是否领取了附件
	ItemData 	[]*ItemDataSt	// 附件
}

// 附件结构体
type ItemDataSt struct {
	ItemType 	int		// 道具类型
	ItemId  	int		// 道具ID
	ItemNum 	int64 	// 道具数量
}
