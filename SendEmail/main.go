package main

import (
	"SendEmail/email"
	"fmt"
)

func main() {
	//新建一个邮件
	newEmail := email.NewEmail("hahayacoder@gmail.com;hahayatest@126.com", "测试golang发送邮件", "这是一个使用golang发送邮件的示例程序.")

	fmt.Println("开始发送邮件...")

	email.SendEmail(newEmail)

	fmt.Println("完成发送邮件...")
}
