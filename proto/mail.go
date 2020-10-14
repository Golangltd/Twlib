package twlib_proto

const (
	// 邮件服务器内部消息协议
	S2CSSendTemplateMailToPlayer = 113 // 发送模板邮件给玩家
)

type S2CSSendTemplateMailToPlayerMsg struct {
	Protocol       int
	Protocol2      int
	OpenID         string // 有值表示发送在线邮件,无值表示发送离线邮件
	UID            int64
	MailTemplateID int
}

func NewS2CSSendTemplateMailToPlayerMsg() *S2CSSendTemplateMailToPlayerMsg {
	m := &S2CSSendTemplateMailToPlayerMsg{
		Protocol:  GGameCenterProto,
		Protocol2: S2CSSendTemplateMailToPlayer,
	}
	return m
}
