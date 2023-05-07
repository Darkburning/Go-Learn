package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func GenVerificationCode(width int) string {
	number := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(number)

	rand.NewSource(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", number[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}

func SendMailVerify(em []string) (string, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("QA注册 <1581690775@qq.com>")
	e.To = em
	// 生成6位随机验证码
	Code := GenVerificationCode(6)
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf("<div>\n\t\t<div>\n\t\t\t尊敬的%s，您好！\n\t\t</div>\n\t\t<div style=\"padding: 8px 40px 8px 50px;\">\n\t\t\t<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>\n\t\t</div>\n\t\t<div>\n\t\t\t<p>此邮箱为系统邮箱，请勿回复。</p>\n\t\t</div>\n\t</div>", em[0], t, Code)
	e.Subject = "QACommunity注册"
	e.HTML = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1581690775@qq.com", "ytglcqbvxyzohcfc", "smtp.qq.com"))
	return Code, err
}

// 发送邮件验证码demo
func main() {
	receiverEmail := make([]string, 1)
	receiverEmail[0] = "1581690775@qq.com"
	verifyCode, err := SendMailVerify(receiverEmail)
	if err != nil {
		fmt.Println(verifyCode)
		os.Exit(1)
	}
}
