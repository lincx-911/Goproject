package model




// MailboxConf 邮箱配置
type MailboxConf struct {
	// 邮件标题
	Title string `json:"title"`
	// 邮件内容
	Body string  `json:"body"`
	// 收件人列表
	RecipientList []string `json:"recipientlist"`
	// 发件人账号
	Sender string   `json:"sender"`
	// 发件人密码，QQ邮箱这里配置授权码
	SPassword string  `json:"spassword"`
	// SMTP 服务器地址， QQ邮箱是smtp.qq.com
	SMTPAddr string  `json:"swtpaddr"`
	// SMTP端口 QQ邮箱是25
	SMTPPort int   `json:"smtpport"`
}

