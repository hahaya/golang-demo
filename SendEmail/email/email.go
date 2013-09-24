package email

import (
	"net/smtp"
	"strings"
)

const (
	HOST        = "smtp.163.com"    //邮件发送者邮箱主机地址
	SERVER_ADDR = "smtp.163.com:25" //邮件发送者邮箱主机服务地址
	USERNAME    = "xxx@163.com"     //邮件发送者邮箱用户名(注意修改成正确的用户名)
	PASSWORD    = "xxx"             //邮件发送者邮箱密码(注意修改成正确的密码)
)

type Email struct {
	to      string //邮件接收者 多个接收者使用;分隔
	subject string //邮件主题
	msg     string //邮件内容
}

//新建邮件函数
func NewEmail(_to, _subject, _msg string) *Email {
	return &Email{to: _to, subject: _subject, msg: _msg}
}

//发送邮件
func SendEmail(email *Email) {
	//连接邮件服务器
	auth := smtp.PlainAuth("", USERNAME, PASSWORD, HOST)
	//使用;分隔多个接收者
	sendTo := strings.Split(email.to, ";")
	//通过从chan中取出值来保证邮件发送到所有的接收者
	doneCh := make(chan error, 1024)

	go func() {
		defer close(doneCh)

		//取出每个邮件接收者
		for _, send := range sendTo {
			//注意邮件格式
			emailStr := "From: " + USERNAME + "\r\nTo:" + send + "\r\nSubject:" + email.subject + "\r\n\r\n" + email.msg

			//发送邮件
			err := smtp.SendMail(SERVER_ADDR, auth, USERNAME, []string{send}, []byte(emailStr))

			//放入chan中
			doneCh <- err
		}
	}()

	//取出chan中的所有值 只有当所以值都能取出时 表示发送到所有接收者
	for i := 0; i < len(sendTo); i++ {
		<-doneCh
	}
}
